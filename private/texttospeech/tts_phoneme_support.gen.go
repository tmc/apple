// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSPhonemeSupport] class.
var (
	_TTSPhonemeSupportClass     TTSPhonemeSupportClass
	_TTSPhonemeSupportClassOnce sync.Once
)

func getTTSPhonemeSupportClass() TTSPhonemeSupportClass {
	_TTSPhonemeSupportClassOnce.Do(func() {
		_TTSPhonemeSupportClass = TTSPhonemeSupportClass{class: objc.GetClass("TTSPhonemeSupport")}
	})
	return _TTSPhonemeSupportClass
}

// GetTTSPhonemeSupportClass returns the class object for TTSPhonemeSupport.
func GetTTSPhonemeSupportClass() TTSPhonemeSupportClass {
	return getTTSPhonemeSupportClass()
}

type TTSPhonemeSupportClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSPhonemeSupportClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSPhonemeSupportClass) Alloc() TTSPhonemeSupport {
	rv := objc.Send[TTSPhonemeSupport](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport
type TTSPhonemeSupport struct {
	objectivec.Object
}

// TTSPhonemeSupportFromID constructs a [TTSPhonemeSupport] from an objc.ID.
func TTSPhonemeSupportFromID(id objc.ID) TTSPhonemeSupport {
	return TTSPhonemeSupport{objectivec.Object{ID: id}}
}

// Ensure TTSPhonemeSupport implements ITTSPhonemeSupport.
var _ ITTSPhonemeSupport = TTSPhonemeSupport{}

// An interface definition for the [TTSPhonemeSupport] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport
type ITTSPhonemeSupport interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSPhonemeSupport) Init() TTSPhonemeSupport {
	rv := objc.Send[TTSPhonemeSupport](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSPhonemeSupport) Autorelease() TTSPhonemeSupport {
	rv := objc.Send[TTSPhonemeSupport](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSPhonemeSupport creates a new TTSPhonemeSupport instance.
func NewTTSPhonemeSupport() TTSPhonemeSupport {
	class := getTTSPhonemeSupportClass()
	rv := objc.Send[TTSPhonemeSupport](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport/_getPhonemeMapForSynth:language:
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) _getPhonemeMapForSynthLanguage(synth objectivec.IObject, language objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSPhonemeSupportClass.class), objc.Sel("_getPhonemeMapForSynth:language:"), synth, language)
	return objectivec.Object{ID: rv}
}

// GetPhonemeMapForSynthLanguage is an exported wrapper for the private method _getPhonemeMapForSynthLanguage.
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) GetPhonemeMapForSynthLanguage(synth objectivec.IObject, language objectivec.IObject) objectivec.IObject {
	return _TTSPhonemeSupportClass._getPhonemeMapForSynthLanguage(synth, language)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport/_ipaVectorFromString:
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) _ipaVectorFromString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSPhonemeSupportClass.class), objc.Sel("_ipaVectorFromString:"), string_)
	return objectivec.Object{ID: rv}
}

// IpaVectorFromString is an exported wrapper for the private method _ipaVectorFromString.
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) IpaVectorFromString(string_ objectivec.IObject) objectivec.IObject {
	return _TTSPhonemeSupportClass._ipaVectorFromString(string_)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport/_phonemesFromIPA:language:synth:
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) _phonemesFromIPALanguageSynth(ipa objectivec.IObject, language objectivec.IObject, synth objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSPhonemeSupportClass.class), objc.Sel("_phonemesFromIPA:language:synth:"), ipa, language, synth)
	return objectivec.Object{ID: rv}
}

// PhonemesFromIPALanguageSynth is an exported wrapper for the private method _phonemesFromIPALanguageSynth.
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) PhonemesFromIPALanguageSynth(ipa objectivec.IObject, language objectivec.IObject, synth objectivec.IObject) objectivec.IObject {
	return _TTSPhonemeSupportClass._phonemesFromIPALanguageSynth(ipa, language, synth)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport/applebetPhonemesFromIPA:
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) ApplebetPhonemesFromIPA(ipa objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSPhonemeSupportClass.class), objc.Sel("applebetPhonemesFromIPA:"), ipa)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport/applebetPhonemesFromLH:
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) ApplebetPhonemesFromLH(lh objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSPhonemeSupportClass.class), objc.Sel("applebetPhonemesFromLH:"), lh)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport/eloquencePhonemesFromIPA:language:
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) EloquencePhonemesFromIPALanguage(ipa objectivec.IObject, language objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSPhonemeSupportClass.class), objc.Sel("eloquencePhonemesFromIPA:language:"), ipa, language)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport/lhPhonemesFromIPA:language:
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) LhPhonemesFromIPALanguage(ipa objectivec.IObject, language objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSPhonemeSupportClass.class), objc.Sel("lhPhonemesFromIPA:language:"), ipa, language)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport/phonemesFromIPA:language:
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) PhonemesFromIPALanguage(ipa objectivec.IObject, language objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSPhonemeSupportClass.class), objc.Sel("phonemesFromIPA:language:"), ipa, language)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSPhonemeSupport/supportedIPAPhonemeLanguages
func (_TTSPhonemeSupportClass TTSPhonemeSupportClass) SupportedIPAPhonemeLanguages() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSPhonemeSupportClass.class), objc.Sel("supportedIPAPhonemeLanguages"))
	return objectivec.Object{ID: rv}
}
