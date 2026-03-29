// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioEngine] class.
var (
	_AVAudioEngineClass     AVAudioEngineClass
	_AVAudioEngineClassOnce sync.Once
)

func getAVAudioEngineClass() AVAudioEngineClass {
	_AVAudioEngineClassOnce.Do(func() {
		_AVAudioEngineClass = AVAudioEngineClass{class: objc.GetClass("AVAudioEngine")}
	})
	return _AVAudioEngineClass
}

// GetAVAudioEngineClass returns the class object for AVAudioEngine.
func GetAVAudioEngineClass() AVAudioEngineClass {
	return getAVAudioEngineClass()
}

type AVAudioEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioEngineClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioEngineClass) Alloc() AVAudioEngine {
	rv := objc.Send[AVAudioEngine](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioEngine.ConnectMIDIToFormatBlock]
//   - [AVAudioEngine.ConnectMIDIToNodesFormatBlock]
//   - [AVAudioEngine.Implementation]
//   - [AVAudioEngine.AutoShutdownEnabled]
//   - [AVAudioEngine.SetAutoShutdownEnabled]
//   - [AVAudioEngine.Running]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine
type AVAudioEngine struct {
	objectivec.Object
}

// AVAudioEngineFromID constructs a [AVAudioEngine] from an objc.ID.
func AVAudioEngineFromID(id objc.ID) AVAudioEngine {
	return AVAudioEngine{objectivec.Object{ID: id}}
}
// Ensure AVAudioEngine implements IAVAudioEngine.
var _ IAVAudioEngine = AVAudioEngine{}

// An interface definition for the [AVAudioEngine] class.
//
// # Methods
//
//   - [IAVAudioEngine.ConnectMIDIToFormatBlock]
//   - [IAVAudioEngine.ConnectMIDIToNodesFormatBlock]
//   - [IAVAudioEngine.Implementation]
//   - [IAVAudioEngine.AutoShutdownEnabled]
//   - [IAVAudioEngine.SetAutoShutdownEnabled]
//   - [IAVAudioEngine.Running]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine
type IAVAudioEngine interface {
	objectivec.IObject

	// Topic: Methods

	ConnectMIDIToFormatBlock(midi objectivec.IObject, to objectivec.IObject, format objectivec.IObject, block VoidHandler)
	ConnectMIDIToNodesFormatBlock(midi objectivec.IObject, nodes objectivec.IObject, format objectivec.IObject, block VoidHandler)
	Implementation() unsafe.Pointer
	AutoShutdownEnabled() bool
	SetAutoShutdownEnabled(value bool)
	Running() bool
}

// Init initializes the instance.
func (a AVAudioEngine) Init() AVAudioEngine {
	rv := objc.Send[AVAudioEngine](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioEngine) Autorelease() AVAudioEngine {
	rv := objc.Send[AVAudioEngine](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioEngine creates a new AVAudioEngine instance.
func NewAVAudioEngine() AVAudioEngine {
	class := getAVAudioEngineClass()
	rv := objc.Send[AVAudioEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connectMIDI:to:format:block:
func (a AVAudioEngine) ConnectMIDIToFormatBlock(midi objectivec.IObject, to objectivec.IObject, format objectivec.IObject, block VoidHandler) {
_block3, _ := NewVoidBlock(block)
	objc.Send[objc.ID](a.ID, objc.Sel("connectMIDI:to:format:block:"), midi, to, format, _block3)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/connectMIDI:toNodes:format:block:
func (a AVAudioEngine) ConnectMIDIToNodesFormatBlock(midi objectivec.IObject, nodes objectivec.IObject, format objectivec.IObject, block VoidHandler) {
_block3, _ := NewVoidBlock(block)
	objc.Send[objc.ID](a.ID, objc.Sel("connectMIDI:toNodes:format:block:"), midi, nodes, format, _block3)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/implementation
func (a AVAudioEngine) Implementation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("implementation"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/autoShutdownEnabled
func (a AVAudioEngine) AutoShutdownEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("autoShutdownEnabled"))
	return rv
}
func (a AVAudioEngine) SetAutoShutdownEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAutoShutdownEnabled:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine/running
func (a AVAudioEngine) Running() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("running"))
	return rv
}

// ConnectMIDIToFormatBlockSync is a synchronous wrapper around [AVAudioEngine.ConnectMIDIToFormatBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioEngine) ConnectMIDIToFormatBlockSync(ctx context.Context, midi objectivec.IObject, to objectivec.IObject, format objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.ConnectMIDIToFormatBlock(midi, to, format, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ConnectMIDIToNodesFormatBlockSync is a synchronous wrapper around [AVAudioEngine.ConnectMIDIToNodesFormatBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioEngine) ConnectMIDIToNodesFormatBlockSync(ctx context.Context, midi objectivec.IObject, nodes objectivec.IObject, format objectivec.IObject) error {
	done := make(chan struct{}, 1)
	a.ConnectMIDIToNodesFormatBlock(midi, nodes, format, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

