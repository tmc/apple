// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAUMessagingAU] class.
var (
	_TTSAUMessagingAUClass     TTSAUMessagingAUClass
	_TTSAUMessagingAUClassOnce sync.Once
)

func getTTSAUMessagingAUClass() TTSAUMessagingAUClass {
	_TTSAUMessagingAUClassOnce.Do(func() {
		_TTSAUMessagingAUClass = TTSAUMessagingAUClass{class: objc.GetClass("TTSAUMessagingAU")}
	})
	return _TTSAUMessagingAUClass
}

// GetTTSAUMessagingAUClass returns the class object for TTSAUMessagingAU.
func GetTTSAUMessagingAUClass() TTSAUMessagingAUClass {
	return getTTSAUMessagingAUClass()
}

type TTSAUMessagingAUClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAUMessagingAUClass) Alloc() TTSAUMessagingAU {
	rv := objc.Send[TTSAUMessagingAU](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSAUMessagingAU.CallAudioUnit]
//   - [TTSAUMessagingAU.Echo]
//   - [TTSAUMessagingAU.OwningAudioUnit]
//   - [TTSAUMessagingAU.SetOwningAudioUnit]
//   - [TTSAUMessagingAU.SetCallHostBlock]
//   - [TTSAUMessagingAU.SetHostBlock]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingAU
type TTSAUMessagingAU struct {
	objectivec.Object
}

// TTSAUMessagingAUFromID constructs a [TTSAUMessagingAU] from an objc.ID.
func TTSAUMessagingAUFromID(id objc.ID) TTSAUMessagingAU {
	return TTSAUMessagingAU{objectivec.Object{ID: id}}
}
// Ensure TTSAUMessagingAU implements ITTSAUMessagingAU.
var _ ITTSAUMessagingAU = TTSAUMessagingAU{}

// An interface definition for the [TTSAUMessagingAU] class.
//
// # Methods
//
//   - [ITTSAUMessagingAU.CallAudioUnit]
//   - [ITTSAUMessagingAU.Echo]
//   - [ITTSAUMessagingAU.OwningAudioUnit]
//   - [ITTSAUMessagingAU.SetOwningAudioUnit]
//   - [ITTSAUMessagingAU.SetCallHostBlock]
//   - [ITTSAUMessagingAU.SetHostBlock]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingAU
type ITTSAUMessagingAU interface {
	objectivec.IObject

	// Topic: Methods

	CallAudioUnit(unit objectivec.IObject) objectivec.IObject
	Echo(echo objectivec.IObject) objectivec.IObject
	OwningAudioUnit() ITTSFirstPartyAudioUnit
	SetOwningAudioUnit(value ITTSFirstPartyAudioUnit)
	SetCallHostBlock(block VoidHandler)
	SetHostBlock(block VoidHandler)
}

// Init initializes the instance.
func (t TTSAUMessagingAU) Init() TTSAUMessagingAU {
	rv := objc.Send[TTSAUMessagingAU](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAUMessagingAU) Autorelease() TTSAUMessagingAU {
	rv := objc.Send[TTSAUMessagingAU](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAUMessagingAU creates a new TTSAUMessagingAU instance.
func NewTTSAUMessagingAU() TTSAUMessagingAU {
	class := getTTSAUMessagingAUClass()
	rv := objc.Send[TTSAUMessagingAU](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingAU/callAudioUnit:
func (t TTSAUMessagingAU) CallAudioUnit(unit objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("callAudioUnit:"), unit)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingAU/echo:
func (t TTSAUMessagingAU) Echo(echo objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("echo:"), echo)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingAU/setCallHostBlock:
func (t TTSAUMessagingAU) SetCallHostBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](t.ID, objc.Sel("setCallHostBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingAU/setHostBlock:
func (t TTSAUMessagingAU) SetHostBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](t.ID, objc.Sel("setHostBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUMessagingAU/owningAudioUnit
func (t TTSAUMessagingAU) OwningAudioUnit() ITTSFirstPartyAudioUnit {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("owningAudioUnit"))
	return TTSFirstPartyAudioUnitFromID(objc.ID(rv))
}
func (t TTSAUMessagingAU) SetOwningAudioUnit(value ITTSFirstPartyAudioUnit) {
	objc.Send[struct{}](t.ID, objc.Sel("setOwningAudioUnit:"), value)
}

// SetCallHostBlockSync is a synchronous wrapper around [TTSAUMessagingAU.SetCallHostBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSAUMessagingAU) SetCallHostBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetCallHostBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetHostBlockSync is a synchronous wrapper around [TTSAUMessagingAU.SetHostBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSAUMessagingAU) SetHostBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetHostBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

