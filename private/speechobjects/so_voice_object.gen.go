// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOVoiceObject] class.
var (
	_SOVoiceObjectClass     SOVoiceObjectClass
	_SOVoiceObjectClassOnce sync.Once
)

func getSOVoiceObjectClass() SOVoiceObjectClass {
	_SOVoiceObjectClassOnce.Do(func() {
		_SOVoiceObjectClass = SOVoiceObjectClass{class: objc.GetClass("SOVoiceObject")}
	})
	return _SOVoiceObjectClass
}

// GetSOVoiceObjectClass returns the class object for SOVoiceObject.
func GetSOVoiceObjectClass() SOVoiceObjectClass {
	return getSOVoiceObjectClass()
}

type SOVoiceObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOVoiceObjectClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOVoiceObjectClass) Alloc() SOVoiceObject {
	rv := objc.Send[SOVoiceObject](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [SOVoiceObject._conversionLocale]
//   - [SOVoiceObject._displayLocalizedVoiceNameForString]
//   - [SOVoiceObject._getSiriVoiceNameFromIdentifier]
//   - [SOVoiceObject._overriddenCompactVoices]
//   - [SOVoiceObject._siriVoiceDisplayName]
//   - [SOVoiceObject._siriVoiceDisplayNameRoot]
//   - [SOVoiceObject._siriVoiceGenderedDisplayName]
//   - [SOVoiceObject._siriVoiceGenderedDisplayNameRoot]
//   - [SOVoiceObject._voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName]
//   - [SOVoiceObject.Compare]
//   - [SOVoiceObject.CountryDisplayString]
//   - [SOVoiceObject.CountryIdentifier]
//   - [SOVoiceObject.DisplayName]
//   - [SOVoiceObject.DisplayNameRoot]
//   - [SOVoiceObject.DoesMatchSystemLocale]
//   - [SOVoiceObject.DownloadPercentComplete]
//   - [SOVoiceObject.SetDownloadPercentComplete]
//   - [SOVoiceObject.DownloadStatus]
//   - [SOVoiceObject.SetDownloadStatus]
//   - [SOVoiceObject.Gender]
//   - [SOVoiceObject.GenderDisplayString]
//   - [SOVoiceObject.Identifier]
//   - [SOVoiceObject.IsAppropriateForSystemLanguage]
//   - [SOVoiceObject.IsNeuter]
//   - [SOVoiceObject.IsSiriVoice]
//   - [SOVoiceObject.LanguageDisplayString]
//   - [SOVoiceObject.LanguageIdentifier]
//   - [SOVoiceObject.LocaleIdentifier]
//   - [SOVoiceObject.MatchesSearchString]
//   - [SOVoiceObject.RelativeDesirability]
//   - [SOVoiceObject.ShowsInFullListOnly]
//   - [SOVoiceObject.SiriLocalizedColorName]
//   - [SOVoiceObject.SiriVoiceDisplayNameFromIdentifier]
//   - [SOVoiceObject.Visibility]
//   - [SOVoiceObject.SetVisibility]
//   - [SOVoiceObject.Voice]
//   - [SOVoiceObject.VoiceAttributes]
//   - [SOVoiceObject.InitWithVoiceWithIdentifier]
//   - [SOVoiceObject.InitWithVoiceIdentifierWithIdentifier]
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject
type SOVoiceObject struct {
	objectivec.Object
}

// SOVoiceObjectFromID constructs a [SOVoiceObject] from an objc.ID.
func SOVoiceObjectFromID(id objc.ID) SOVoiceObject {
	return SOVoiceObject{objectivec.Object{ID: id}}
}
// Ensure SOVoiceObject implements ISOVoiceObject.
var _ ISOVoiceObject = SOVoiceObject{}

// An interface definition for the [SOVoiceObject] class.
//
// # Methods
//
//   - [ISOVoiceObject._conversionLocale]
//   - [ISOVoiceObject._displayLocalizedVoiceNameForString]
//   - [ISOVoiceObject._getSiriVoiceNameFromIdentifier]
//   - [ISOVoiceObject._overriddenCompactVoices]
//   - [ISOVoiceObject._siriVoiceDisplayName]
//   - [ISOVoiceObject._siriVoiceDisplayNameRoot]
//   - [ISOVoiceObject._siriVoiceGenderedDisplayName]
//   - [ISOVoiceObject._siriVoiceGenderedDisplayNameRoot]
//   - [ISOVoiceObject._voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName]
//   - [ISOVoiceObject.Compare]
//   - [ISOVoiceObject.CountryDisplayString]
//   - [ISOVoiceObject.CountryIdentifier]
//   - [ISOVoiceObject.DisplayName]
//   - [ISOVoiceObject.DisplayNameRoot]
//   - [ISOVoiceObject.DoesMatchSystemLocale]
//   - [ISOVoiceObject.DownloadPercentComplete]
//   - [ISOVoiceObject.SetDownloadPercentComplete]
//   - [ISOVoiceObject.DownloadStatus]
//   - [ISOVoiceObject.SetDownloadStatus]
//   - [ISOVoiceObject.Gender]
//   - [ISOVoiceObject.GenderDisplayString]
//   - [ISOVoiceObject.Identifier]
//   - [ISOVoiceObject.IsAppropriateForSystemLanguage]
//   - [ISOVoiceObject.IsNeuter]
//   - [ISOVoiceObject.IsSiriVoice]
//   - [ISOVoiceObject.LanguageDisplayString]
//   - [ISOVoiceObject.LanguageIdentifier]
//   - [ISOVoiceObject.LocaleIdentifier]
//   - [ISOVoiceObject.MatchesSearchString]
//   - [ISOVoiceObject.RelativeDesirability]
//   - [ISOVoiceObject.ShowsInFullListOnly]
//   - [ISOVoiceObject.SiriLocalizedColorName]
//   - [ISOVoiceObject.SiriVoiceDisplayNameFromIdentifier]
//   - [ISOVoiceObject.Visibility]
//   - [ISOVoiceObject.SetVisibility]
//   - [ISOVoiceObject.Voice]
//   - [ISOVoiceObject.VoiceAttributes]
//   - [ISOVoiceObject.InitWithVoiceWithIdentifier]
//   - [ISOVoiceObject.InitWithVoiceIdentifierWithIdentifier]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject
type ISOVoiceObject interface {
	objectivec.IObject

	// Topic: Methods

	_conversionLocale() objectivec.IObject
	_displayLocalizedVoiceNameForString(string_ objectivec.IObject) objectivec.IObject
	_getSiriVoiceNameFromIdentifier() objectivec.IObject
	_overriddenCompactVoices() objectivec.IObject
	_siriVoiceDisplayName() objectivec.IObject
	_siriVoiceDisplayNameRoot() objectivec.IObject
	_siriVoiceGenderedDisplayName() objectivec.IObject
	_siriVoiceGenderedDisplayNameRoot() objectivec.IObject
	_voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName(name objectivec.IObject) objectivec.IObject
	Compare(compare objectivec.IObject) int64
	CountryDisplayString() objectivec.IObject
	CountryIdentifier() objectivec.IObject
	DisplayName() objectivec.IObject
	DisplayNameRoot() objectivec.IObject
	DoesMatchSystemLocale() bool
	DownloadPercentComplete() float64
	SetDownloadPercentComplete(value float64)
	DownloadStatus() uint64
	SetDownloadStatus(value uint64)
	Gender() objectivec.IObject
	GenderDisplayString() objectivec.IObject
	Identifier() objectivec.IObject
	IsAppropriateForSystemLanguage() bool
	IsNeuter() bool
	IsSiriVoice() bool
	LanguageDisplayString() objectivec.IObject
	LanguageIdentifier() objectivec.IObject
	LocaleIdentifier() objectivec.IObject
	MatchesSearchString(string_ objectivec.IObject) bool
	RelativeDesirability() int64
	ShowsInFullListOnly() bool
	SiriLocalizedColorName() objectivec.IObject
	SiriVoiceDisplayNameFromIdentifier() objectivec.IObject
	Visibility() uint64
	SetVisibility(value uint64)
	Voice() unsafe.Pointer
	VoiceAttributes() objectivec.IObject
	InitWithVoiceWithIdentifier(voice objectivec.IObject, identifier objectivec.IObject) SOVoiceObject
	InitWithVoiceIdentifierWithIdentifier(identifier objectivec.IObject) SOVoiceObject
}

// Init initializes the instance.
func (s SOVoiceObject) Init() SOVoiceObject {
	rv := objc.Send[SOVoiceObject](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOVoiceObject) Autorelease() SOVoiceObject {
	rv := objc.Send[SOVoiceObject](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOVoiceObject creates a new SOVoiceObject instance.
func NewSOVoiceObject() SOVoiceObject {
	class := getSOVoiceObjectClass()
	rv := objc.Send[SOVoiceObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/initWithVoice:identifier:
func NewSOVoiceObjectWithVoiceIdentifier(voice objectivec.IObject, identifier objectivec.IObject) SOVoiceObject {
	instance := getSOVoiceObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVoice:identifier:"), voice, identifier)
	return SOVoiceObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/_conversionLocale
func (s SOVoiceObject) _conversionLocale() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_conversionLocale"))
	return objectivec.Object{ID: rv}
}

// ConversionLocale is an exported wrapper for the private method _conversionLocale.
func (s SOVoiceObject) ConversionLocale() objectivec.IObject {
	return s._conversionLocale()
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/_displayLocalizedVoiceNameForString:
func (s SOVoiceObject) _displayLocalizedVoiceNameForString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_displayLocalizedVoiceNameForString:"), string_)
	return objectivec.Object{ID: rv}
}

// DisplayLocalizedVoiceNameForString is an exported wrapper for the private method _displayLocalizedVoiceNameForString.
func (s SOVoiceObject) DisplayLocalizedVoiceNameForString(string_ objectivec.IObject) objectivec.IObject {
	return s._displayLocalizedVoiceNameForString(string_)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/_getSiriVoiceNameFromIdentifier
func (s SOVoiceObject) _getSiriVoiceNameFromIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_getSiriVoiceNameFromIdentifier"))
	return objectivec.Object{ID: rv}
}

// GetSiriVoiceNameFromIdentifier is an exported wrapper for the private method _getSiriVoiceNameFromIdentifier.
func (s SOVoiceObject) GetSiriVoiceNameFromIdentifier() objectivec.IObject {
	return s._getSiriVoiceNameFromIdentifier()
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/_overriddenCompactVoices
func (s SOVoiceObject) _overriddenCompactVoices() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_overriddenCompactVoices"))
	return objectivec.Object{ID: rv}
}

// OverriddenCompactVoices is an exported wrapper for the private method _overriddenCompactVoices.
func (s SOVoiceObject) OverriddenCompactVoices() objectivec.IObject {
	return s._overriddenCompactVoices()
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/_siriVoiceDisplayName
func (s SOVoiceObject) _siriVoiceDisplayName() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_siriVoiceDisplayName"))
	return objectivec.Object{ID: rv}
}

// SiriVoiceDisplayName is an exported wrapper for the private method _siriVoiceDisplayName.
func (s SOVoiceObject) SiriVoiceDisplayName() objectivec.IObject {
	return s._siriVoiceDisplayName()
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/_siriVoiceDisplayNameRoot
func (s SOVoiceObject) _siriVoiceDisplayNameRoot() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_siriVoiceDisplayNameRoot"))
	return objectivec.Object{ID: rv}
}

// SiriVoiceDisplayNameRoot is an exported wrapper for the private method _siriVoiceDisplayNameRoot.
func (s SOVoiceObject) SiriVoiceDisplayNameRoot() objectivec.IObject {
	return s._siriVoiceDisplayNameRoot()
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/_siriVoiceGenderedDisplayName
func (s SOVoiceObject) _siriVoiceGenderedDisplayName() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_siriVoiceGenderedDisplayName"))
	return objectivec.Object{ID: rv}
}

// SiriVoiceGenderedDisplayName is an exported wrapper for the private method _siriVoiceGenderedDisplayName.
func (s SOVoiceObject) SiriVoiceGenderedDisplayName() objectivec.IObject {
	return s._siriVoiceGenderedDisplayName()
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/_siriVoiceGenderedDisplayNameRoot
func (s SOVoiceObject) _siriVoiceGenderedDisplayNameRoot() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_siriVoiceGenderedDisplayNameRoot"))
	return objectivec.Object{ID: rv}
}

// SiriVoiceGenderedDisplayNameRoot is an exported wrapper for the private method _siriVoiceGenderedDisplayNameRoot.
func (s SOVoiceObject) SiriVoiceGenderedDisplayNameRoot() objectivec.IObject {
	return s._siriVoiceGenderedDisplayNameRoot()
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/_voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName:
func (s SOVoiceObject) _voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName:"), name)
	return objectivec.Object{ID: rv}
}

// VoiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName is an exported wrapper for the private method _voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName.
func (s SOVoiceObject) VoiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName(name objectivec.IObject) objectivec.IObject {
	return s._voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName(name)
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/compare:
func (s SOVoiceObject) Compare(compare objectivec.IObject) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("compare:"), compare)
	return rv
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/countryDisplayString
func (s SOVoiceObject) CountryDisplayString() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("countryDisplayString"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/countryIdentifier
func (s SOVoiceObject) CountryIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("countryIdentifier"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/displayName
func (s SOVoiceObject) DisplayName() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayName"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/displayNameRoot
func (s SOVoiceObject) DisplayNameRoot() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayNameRoot"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/doesMatchSystemLocale
func (s SOVoiceObject) DoesMatchSystemLocale() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("doesMatchSystemLocale"))
	return rv
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/gender
func (s SOVoiceObject) Gender() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("gender"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/genderDisplayString
func (s SOVoiceObject) GenderDisplayString() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("genderDisplayString"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/identifier
func (s SOVoiceObject) Identifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("identifier"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/isAppropriateForSystemLanguage
func (s SOVoiceObject) IsAppropriateForSystemLanguage() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isAppropriateForSystemLanguage"))
	return rv
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/isNeuter
func (s SOVoiceObject) IsNeuter() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isNeuter"))
	return rv
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/isSiriVoice
func (s SOVoiceObject) IsSiriVoice() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSiriVoice"))
	return rv
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/languageDisplayString
func (s SOVoiceObject) LanguageDisplayString() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("languageDisplayString"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/languageIdentifier
func (s SOVoiceObject) LanguageIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("languageIdentifier"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/localeIdentifier
func (s SOVoiceObject) LocaleIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("localeIdentifier"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/matchesSearchString:
func (s SOVoiceObject) MatchesSearchString(string_ objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("matchesSearchString:"), string_)
	return rv
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/relativeDesirability
func (s SOVoiceObject) RelativeDesirability() int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("relativeDesirability"))
	return rv
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/showsInFullListOnly
func (s SOVoiceObject) ShowsInFullListOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showsInFullListOnly"))
	return rv
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/siriLocalizedColorName
func (s SOVoiceObject) SiriLocalizedColorName() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("siriLocalizedColorName"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/siriVoiceDisplayNameFromIdentifier
func (s SOVoiceObject) SiriVoiceDisplayNameFromIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("siriVoiceDisplayNameFromIdentifier"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/voiceAttributes
func (s SOVoiceObject) VoiceAttributes() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voiceAttributes"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/initWithVoice:identifier:
func (s SOVoiceObject) InitWithVoiceWithIdentifier(voice objectivec.IObject, identifier objectivec.IObject) SOVoiceObject {
	rv := objc.Send[SOVoiceObject](s.ID, objc.Sel("initWithVoice:identifier:"), voice, identifier)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/initWithVoiceIdentifier:
func (s SOVoiceObject) InitWithVoiceIdentifierWithIdentifier(identifier objectivec.IObject) SOVoiceObject {
	rv := objc.Send[SOVoiceObject](s.ID, objc.Sel("initWithVoiceIdentifier:"), identifier)
	return rv
}

//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/assetForVoiceID:
func (_SOVoiceObjectClass SOVoiceObjectClass) AssetForVoiceID(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("assetForVoiceID:"), id)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/invalidateAssetMaps
func (_SOVoiceObjectClass SOVoiceObjectClass) InvalidateAssetMaps() {
	objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("invalidateAssetMaps"))
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/isSameLanguageFromLocaleIdentifier:secondLocaleIdentifier:
func (_SOVoiceObjectClass SOVoiceObjectClass) IsSameLanguageFromLocaleIdentifierSecondLocaleIdentifier(identifier objectivec.IObject, identifier2 objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SOVoiceObjectClass.class), objc.Sel("isSameLanguageFromLocaleIdentifier:secondLocaleIdentifier:"), identifier, identifier2)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/isSameLocaleIdentifier:secondLocaleIdentifier:
func (_SOVoiceObjectClass SOVoiceObjectClass) IsSameLocaleIdentifierSecondLocaleIdentifier(identifier objectivec.IObject, identifier2 objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SOVoiceObjectClass.class), objc.Sel("isSameLocaleIdentifier:secondLocaleIdentifier:"), identifier, identifier2)
	return rv
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/normalizedVoiceIdentifier:
func (_SOVoiceObjectClass SOVoiceObjectClass) NormalizedVoiceIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("normalizedVoiceIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/rebuildAssetMaps
func (_SOVoiceObjectClass SOVoiceObjectClass) RebuildAssetMaps() {
	objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("rebuildAssetMaps"))
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/setVisibleVoicesTable:
func (_SOVoiceObjectClass SOVoiceObjectClass) SetVisibleVoicesTable(table objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("setVisibleVoicesTable:"), table)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/systemLocaleIdentifier
func (_SOVoiceObjectClass SOVoiceObjectClass) SystemLocaleIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("systemLocaleIdentifier"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/visibleVoicesForLocaleIdentifier:additionalRequiredVoices:allowAllVoices:
func (_SOVoiceObjectClass SOVoiceObjectClass) VisibleVoicesForLocaleIdentifierAdditionalRequiredVoicesAllowAllVoices(identifier objectivec.IObject, voices objectivec.IObject, voices2 bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("visibleVoicesForLocaleIdentifier:additionalRequiredVoices:allowAllVoices:"), identifier, voices, voices2)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/visibleVoicesTable
func (_SOVoiceObjectClass SOVoiceObjectClass) VisibleVoicesTable() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("visibleVoicesTable"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/downloadPercentComplete
func (s SOVoiceObject) DownloadPercentComplete() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("downloadPercentComplete"))
	return rv
}
func (s SOVoiceObject) SetDownloadPercentComplete(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setDownloadPercentComplete:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/downloadStatus
func (s SOVoiceObject) DownloadStatus() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("downloadStatus"))
	return rv
}
func (s SOVoiceObject) SetDownloadStatus(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setDownloadStatus:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/visibility
func (s SOVoiceObject) Visibility() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("visibility"))
	return rv
}
func (s SOVoiceObject) SetVisibility(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setVisibility:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOVoiceObject/voice
func (s SOVoiceObject) Voice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("voice"))
	return rv
}

