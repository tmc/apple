// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitComponent] class.
var (
	_AVAudioUnitComponentClass     AVAudioUnitComponentClass
	_AVAudioUnitComponentClassOnce sync.Once
)

func getAVAudioUnitComponentClass() AVAudioUnitComponentClass {
	_AVAudioUnitComponentClassOnce.Do(func() {
		_AVAudioUnitComponentClass = AVAudioUnitComponentClass{class: objc.GetClass("AVAudioUnitComponent")}
	})
	return _AVAudioUnitComponentClass
}

// GetAVAudioUnitComponentClass returns the class object for AVAudioUnitComponent.
func GetAVAudioUnitComponentClass() AVAudioUnitComponentClass {
	return getAVAudioUnitComponentClass()
}

type AVAudioUnitComponentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitComponentClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitComponentClass) Alloc() AVAudioUnitComponent {
	rv := objc.Send[AVAudioUnitComponent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioUnitComponent.GetTypeName]
//   - [AVAudioUnitComponent.ComponentURL]
//   - [AVAudioUnitComponent.LocaleChanged]
//   - [AVAudioUnitComponent.ValidateWithResultsInCompletionHandler]
//   - [AVAudioUnitComponent.SandboxSafe]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent
type AVAudioUnitComponent struct {
	objectivec.Object
}

// AVAudioUnitComponentFromID constructs a [AVAudioUnitComponent] from an objc.ID.
func AVAudioUnitComponentFromID(id objc.ID) AVAudioUnitComponent {
	return AVAudioUnitComponent{objectivec.Object{ID: id}}
}

// Ensure AVAudioUnitComponent implements IAVAudioUnitComponent.
var _ IAVAudioUnitComponent = AVAudioUnitComponent{}

// An interface definition for the [AVAudioUnitComponent] class.
//
// # Methods
//
//   - [IAVAudioUnitComponent.GetTypeName]
//   - [IAVAudioUnitComponent.ComponentURL]
//   - [IAVAudioUnitComponent.LocaleChanged]
//   - [IAVAudioUnitComponent.ValidateWithResultsInCompletionHandler]
//   - [IAVAudioUnitComponent.SandboxSafe]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent
type IAVAudioUnitComponent interface {
	objectivec.IObject

	// Topic: Methods

	GetTypeName(name uint32) objectivec.IObject
	ComponentURL() foundation.INSURL
	LocaleChanged()
	ValidateWithResultsInCompletionHandler(results objectivec.IObject, handler ErrorHandler) bool
	SandboxSafe() bool
}

// Init initializes the instance.
func (a AVAudioUnitComponent) Init() AVAudioUnitComponent {
	rv := objc.Send[AVAudioUnitComponent](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitComponent) Autorelease() AVAudioUnitComponent {
	rv := objc.Send[AVAudioUnitComponent](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitComponent creates a new AVAudioUnitComponent instance.
func NewAVAudioUnitComponent() AVAudioUnitComponent {
	class := getAVAudioUnitComponentClass()
	rv := objc.Send[AVAudioUnitComponent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/GetTypeName:
func (a AVAudioUnitComponent) GetTypeName(name uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("GetTypeName:"), name)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/localeChanged
func (a AVAudioUnitComponent) LocaleChanged() {
	objc.Send[objc.ID](a.ID, objc.Sel("localeChanged"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/validateWithResults:inCompletionHandler:
func (a AVAudioUnitComponent) ValidateWithResultsInCompletionHandler(results objectivec.IObject, handler ErrorHandler) bool {
	_block1, _ := NewErrorBlock(handler)
	rv := objc.Send[bool](a.ID, objc.Sel("validateWithResults:inCompletionHandler:"), results, _block1)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/componentURL
func (a AVAudioUnitComponent) ComponentURL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("componentURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponent/sandboxSafe
func (a AVAudioUnitComponent) SandboxSafe() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("sandboxSafe"))
	return rv
}

// ValidateWithResultsIn is a synchronous wrapper around [AVAudioUnitComponent.ValidateWithResultsInCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioUnitComponent) ValidateWithResultsIn(ctx context.Context, results objectivec.IObject) error {
	done := make(chan error, 1)
	a.ValidateWithResultsInCompletionHandler(results, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
