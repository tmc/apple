// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAlternativeVoices] class.
var (
	_TTSAlternativeVoicesClass     TTSAlternativeVoicesClass
	_TTSAlternativeVoicesClassOnce sync.Once
)

func getTTSAlternativeVoicesClass() TTSAlternativeVoicesClass {
	_TTSAlternativeVoicesClassOnce.Do(func() {
		_TTSAlternativeVoicesClass = TTSAlternativeVoicesClass{class: objc.GetClass("TTSAlternativeVoices")}
	})
	return _TTSAlternativeVoicesClass
}

// GetTTSAlternativeVoicesClass returns the class object for TTSAlternativeVoices.
func GetTTSAlternativeVoicesClass() TTSAlternativeVoicesClass {
	return getTTSAlternativeVoicesClass()
}

type TTSAlternativeVoicesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSAlternativeVoicesClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAlternativeVoicesClass) Alloc() TTSAlternativeVoices {
	rv := objc.Send[TTSAlternativeVoices](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices
type TTSAlternativeVoices struct {
	objectivec.Object
}

// TTSAlternativeVoicesFromID constructs a [TTSAlternativeVoices] from an objc.ID.
func TTSAlternativeVoicesFromID(id objc.ID) TTSAlternativeVoices {
	return TTSAlternativeVoices{objectivec.Object{ID: id}}
}

// Ensure TTSAlternativeVoices implements ITTSAlternativeVoices.
var _ ITTSAlternativeVoices = TTSAlternativeVoices{}

// An interface definition for the [TTSAlternativeVoices] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices
type ITTSAlternativeVoices interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSAlternativeVoices) Init() TTSAlternativeVoices {
	rv := objc.Send[TTSAlternativeVoices](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAlternativeVoices) Autorelease() TTSAlternativeVoices {
	rv := objc.Send[TTSAlternativeVoices](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAlternativeVoices creates a new TTSAlternativeVoices instance.
func NewTTSAlternativeVoices() TTSAlternativeVoices {
	class := getTTSAlternativeVoicesClass()
	rv := objc.Send[TTSAlternativeVoices](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices/isCombinedVocalizerVoiceIdentifier:
func (_TTSAlternativeVoicesClass TTSAlternativeVoicesClass) IsCombinedVocalizerVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSAlternativeVoicesClass.class), objc.Sel("isCombinedVocalizerVoiceIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices/isCompactVocalizerVoiceIdentifier:
func (_TTSAlternativeVoicesClass TTSAlternativeVoicesClass) IsCompactVocalizerVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSAlternativeVoicesClass.class), objc.Sel("isCompactVocalizerVoiceIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices/isNeuralAXSiriVoiceIdentifier:
func (_TTSAlternativeVoicesClass TTSAlternativeVoicesClass) IsNeuralAXSiriVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSAlternativeVoicesClass.class), objc.Sel("isNeuralAXSiriVoiceIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices/isNeuralSiriVoiceIdentifier:
func (_TTSAlternativeVoicesClass TTSAlternativeVoicesClass) IsNeuralSiriVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSAlternativeVoicesClass.class), objc.Sel("isNeuralSiriVoiceIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices/isOldSiriVoiceIdentifier:
func (_TTSAlternativeVoicesClass TTSAlternativeVoicesClass) IsOldSiriVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSAlternativeVoicesClass.class), objc.Sel("isOldSiriVoiceIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices/isPersonalVoiceIdentifier:
func (_TTSAlternativeVoicesClass TTSAlternativeVoicesClass) IsPersonalVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSAlternativeVoicesClass.class), objc.Sel("isPersonalVoiceIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices/isSiriVoiceIdentifier:
func (_TTSAlternativeVoicesClass TTSAlternativeVoicesClass) IsSiriVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSAlternativeVoicesClass.class), objc.Sel("isSiriVoiceIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices/isVocalizerVoiceIdentifier:
func (_TTSAlternativeVoicesClass TTSAlternativeVoicesClass) IsVocalizerVoiceIdentifier(identifier objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSAlternativeVoicesClass.class), objc.Sel("isVocalizerVoiceIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAlternativeVoices/nameForVoiceIdentifier:
func (_TTSAlternativeVoicesClass TTSAlternativeVoicesClass) NameForVoiceIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSAlternativeVoicesClass.class), objc.Sel("nameForVoiceIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
