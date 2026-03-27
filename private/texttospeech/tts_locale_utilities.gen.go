// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSLocaleUtilities] class.
var (
	_TTSLocaleUtilitiesClass     TTSLocaleUtilitiesClass
	_TTSLocaleUtilitiesClassOnce sync.Once
)

func getTTSLocaleUtilitiesClass() TTSLocaleUtilitiesClass {
	_TTSLocaleUtilitiesClassOnce.Do(func() {
		_TTSLocaleUtilitiesClass = TTSLocaleUtilitiesClass{class: objc.GetClass("TTSLocaleUtilities")}
	})
	return _TTSLocaleUtilitiesClass
}

// GetTTSLocaleUtilitiesClass returns the class object for TTSLocaleUtilities.
func GetTTSLocaleUtilitiesClass() TTSLocaleUtilitiesClass {
	return getTTSLocaleUtilitiesClass()
}

type TTSLocaleUtilitiesClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSLocaleUtilitiesClass) Alloc() TTSLocaleUtilities {
	rv := objc.Send[TTSLocaleUtilities](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSLocaleUtilities.CanonicalLanguageCodeVoiceNamesData]
//   - [TTSLocaleUtilities.SetCanonicalLanguageCodeVoiceNamesData]
//   - [TTSLocaleUtilities.DefaultVoiceIdentifierForGeneralLanguageCode]
//   - [TTSLocaleUtilities.DefaultVoiceIdentifierForVoiceName]
//   - [TTSLocaleUtilities.FallbackSampleStringCache]
//   - [TTSLocaleUtilities.SetFallbackSampleStringCache]
//   - [TTSLocaleUtilities.GeneralLanguageCodeData]
//   - [TTSLocaleUtilities.SetGeneralLanguageCodeData]
//   - [TTSLocaleUtilities.SampleStringForVoiceIdentifier]
//   - [TTSLocaleUtilities.SampleStringForVoiceIdentifierWithPreferredLocaleID]
//   - [TTSLocaleUtilities.VoiceIdSampleStringData]
//   - [TTSLocaleUtilities.SetVoiceIdSampleStringData]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities
type TTSLocaleUtilities struct {
	objectivec.Object
}

// TTSLocaleUtilitiesFromID constructs a [TTSLocaleUtilities] from an objc.ID.
func TTSLocaleUtilitiesFromID(id objc.ID) TTSLocaleUtilities {
	return TTSLocaleUtilities{objectivec.Object{ID: id}}
}
// Ensure TTSLocaleUtilities implements ITTSLocaleUtilities.
var _ ITTSLocaleUtilities = TTSLocaleUtilities{}

// An interface definition for the [TTSLocaleUtilities] class.
//
// # Methods
//
//   - [ITTSLocaleUtilities.CanonicalLanguageCodeVoiceNamesData]
//   - [ITTSLocaleUtilities.SetCanonicalLanguageCodeVoiceNamesData]
//   - [ITTSLocaleUtilities.DefaultVoiceIdentifierForGeneralLanguageCode]
//   - [ITTSLocaleUtilities.DefaultVoiceIdentifierForVoiceName]
//   - [ITTSLocaleUtilities.FallbackSampleStringCache]
//   - [ITTSLocaleUtilities.SetFallbackSampleStringCache]
//   - [ITTSLocaleUtilities.GeneralLanguageCodeData]
//   - [ITTSLocaleUtilities.SetGeneralLanguageCodeData]
//   - [ITTSLocaleUtilities.SampleStringForVoiceIdentifier]
//   - [ITTSLocaleUtilities.SampleStringForVoiceIdentifierWithPreferredLocaleID]
//   - [ITTSLocaleUtilities.VoiceIdSampleStringData]
//   - [ITTSLocaleUtilities.SetVoiceIdSampleStringData]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities
type ITTSLocaleUtilities interface {
	objectivec.IObject

	// Topic: Methods

	CanonicalLanguageCodeVoiceNamesData() foundation.INSDictionary
	SetCanonicalLanguageCodeVoiceNamesData(value foundation.INSDictionary)
	DefaultVoiceIdentifierForGeneralLanguageCode(code objectivec.IObject) objectivec.IObject
	DefaultVoiceIdentifierForVoiceName(name objectivec.IObject) objectivec.IObject
	FallbackSampleStringCache() foundation.INSDictionary
	SetFallbackSampleStringCache(value foundation.INSDictionary)
	GeneralLanguageCodeData() foundation.INSDictionary
	SetGeneralLanguageCodeData(value foundation.INSDictionary)
	SampleStringForVoiceIdentifier(identifier objectivec.IObject) objectivec.IObject
	SampleStringForVoiceIdentifierWithPreferredLocaleID(identifier objectivec.IObject, id objectivec.IObject) objectivec.IObject
	VoiceIdSampleStringData() foundation.INSDictionary
	SetVoiceIdSampleStringData(value foundation.INSDictionary)
}

// Init initializes the instance.
func (t TTSLocaleUtilities) Init() TTSLocaleUtilities {
	rv := objc.Send[TTSLocaleUtilities](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSLocaleUtilities) Autorelease() TTSLocaleUtilities {
	rv := objc.Send[TTSLocaleUtilities](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSLocaleUtilities creates a new TTSLocaleUtilities instance.
func NewTTSLocaleUtilities() TTSLocaleUtilities {
	class := getTTSLocaleUtilitiesClass()
	rv := objc.Send[TTSLocaleUtilities](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities/defaultVoiceIdentifierForGeneralLanguageCode:
func (t TTSLocaleUtilities) DefaultVoiceIdentifierForGeneralLanguageCode(code objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("defaultVoiceIdentifierForGeneralLanguageCode:"), code)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities/defaultVoiceIdentifierForVoiceName:
func (t TTSLocaleUtilities) DefaultVoiceIdentifierForVoiceName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("defaultVoiceIdentifierForVoiceName:"), name)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities/sampleStringForVoiceIdentifier:
func (t TTSLocaleUtilities) SampleStringForVoiceIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("sampleStringForVoiceIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities/sampleStringForVoiceIdentifier:withPreferredLocaleID:
func (t TTSLocaleUtilities) SampleStringForVoiceIdentifierWithPreferredLocaleID(identifier objectivec.IObject, id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("sampleStringForVoiceIdentifier:withPreferredLocaleID:"), identifier, id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities/sharedInstance
func (_TTSLocaleUtilitiesClass TTSLocaleUtilitiesClass) SharedInstance() TTSLocaleUtilities {
	rv := objc.Send[objc.ID](objc.ID(_TTSLocaleUtilitiesClass.class), objc.Sel("sharedInstance"))
	return TTSLocaleUtilitiesFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities/canonicalLanguageCodeVoiceNamesData
func (t TTSLocaleUtilities) CanonicalLanguageCodeVoiceNamesData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("canonicalLanguageCodeVoiceNamesData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSLocaleUtilities) SetCanonicalLanguageCodeVoiceNamesData(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setCanonicalLanguageCodeVoiceNamesData:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities/fallbackSampleStringCache
func (t TTSLocaleUtilities) FallbackSampleStringCache() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("fallbackSampleStringCache"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSLocaleUtilities) SetFallbackSampleStringCache(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setFallbackSampleStringCache:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities/generalLanguageCodeData
func (t TTSLocaleUtilities) GeneralLanguageCodeData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("generalLanguageCodeData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSLocaleUtilities) SetGeneralLanguageCodeData(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setGeneralLanguageCodeData:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLocaleUtilities/voiceIdSampleStringData
func (t TTSLocaleUtilities) VoiceIdSampleStringData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceIdSampleStringData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSLocaleUtilities) SetVoiceIdSampleStringData(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoiceIdSampleStringData:"), value)
}

