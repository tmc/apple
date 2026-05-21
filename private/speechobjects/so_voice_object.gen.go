// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"unsafe"

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

func NewSOVoiceObjectWithVoiceIdentifier(voice objectivec.IObject, identifier objectivec.IObject) SOVoiceObject {
	instance := getSOVoiceObjectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVoice:identifier:"), voice, identifier)
	return SOVoiceObjectFromID(rv)
}

func (s SOVoiceObject) _conversionLocale() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_conversionLocale"))
	return objectivec.Object{ID: rv}
}

// ConversionLocale is an exported wrapper for the private method _conversionLocale.
func (s SOVoiceObject) ConversionLocale() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_conversionLocale")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_conversionLocale"}
		return nil, err
	}
	return s._conversionLocale(), nil
}

// CanConversionLocale reports whether the receiver responds to the private selector _conversionLocale.
func (s SOVoiceObject) CanConversionLocale() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_conversionLocale"))
}
func (s SOVoiceObject) _displayLocalizedVoiceNameForString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_displayLocalizedVoiceNameForString:"), string_)
	return objectivec.Object{ID: rv}
}

// DisplayLocalizedVoiceNameForString is an exported wrapper for the private method _displayLocalizedVoiceNameForString.
func (s SOVoiceObject) DisplayLocalizedVoiceNameForString(string_ objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_displayLocalizedVoiceNameForString:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_displayLocalizedVoiceNameForString:"}
		return nil, err
	}
	return s._displayLocalizedVoiceNameForString(string_), nil
}

// CanDisplayLocalizedVoiceNameForString reports whether the receiver responds to the private selector _displayLocalizedVoiceNameForString:.
func (s SOVoiceObject) CanDisplayLocalizedVoiceNameForString() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_displayLocalizedVoiceNameForString:"))
}
func (s SOVoiceObject) _getSiriVoiceNameFromIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_getSiriVoiceNameFromIdentifier"))
	return objectivec.Object{ID: rv}
}

// GetSiriVoiceNameFromIdentifier is an exported wrapper for the private method _getSiriVoiceNameFromIdentifier.
func (s SOVoiceObject) GetSiriVoiceNameFromIdentifier() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_getSiriVoiceNameFromIdentifier")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_getSiriVoiceNameFromIdentifier"}
		return nil, err
	}
	return s._getSiriVoiceNameFromIdentifier(), nil
}

// CanGetSiriVoiceNameFromIdentifier reports whether the receiver responds to the private selector _getSiriVoiceNameFromIdentifier.
func (s SOVoiceObject) CanGetSiriVoiceNameFromIdentifier() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_getSiriVoiceNameFromIdentifier"))
}
func (s SOVoiceObject) _overriddenCompactVoices() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_overriddenCompactVoices"))
	return objectivec.Object{ID: rv}
}

// OverriddenCompactVoices is an exported wrapper for the private method _overriddenCompactVoices.
func (s SOVoiceObject) OverriddenCompactVoices() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_overriddenCompactVoices")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_overriddenCompactVoices"}
		return nil, err
	}
	return s._overriddenCompactVoices(), nil
}

// CanOverriddenCompactVoices reports whether the receiver responds to the private selector _overriddenCompactVoices.
func (s SOVoiceObject) CanOverriddenCompactVoices() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_overriddenCompactVoices"))
}
func (s SOVoiceObject) _siriVoiceDisplayName() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_siriVoiceDisplayName"))
	return objectivec.Object{ID: rv}
}

// SiriVoiceDisplayName is an exported wrapper for the private method _siriVoiceDisplayName.
func (s SOVoiceObject) SiriVoiceDisplayName() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_siriVoiceDisplayName")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_siriVoiceDisplayName"}
		return nil, err
	}
	return s._siriVoiceDisplayName(), nil
}

// CanSiriVoiceDisplayName reports whether the receiver responds to the private selector _siriVoiceDisplayName.
func (s SOVoiceObject) CanSiriVoiceDisplayName() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_siriVoiceDisplayName"))
}
func (s SOVoiceObject) _siriVoiceDisplayNameRoot() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_siriVoiceDisplayNameRoot"))
	return objectivec.Object{ID: rv}
}

// SiriVoiceDisplayNameRoot is an exported wrapper for the private method _siriVoiceDisplayNameRoot.
func (s SOVoiceObject) SiriVoiceDisplayNameRoot() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_siriVoiceDisplayNameRoot")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_siriVoiceDisplayNameRoot"}
		return nil, err
	}
	return s._siriVoiceDisplayNameRoot(), nil
}

// CanSiriVoiceDisplayNameRoot reports whether the receiver responds to the private selector _siriVoiceDisplayNameRoot.
func (s SOVoiceObject) CanSiriVoiceDisplayNameRoot() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_siriVoiceDisplayNameRoot"))
}
func (s SOVoiceObject) _siriVoiceGenderedDisplayName() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_siriVoiceGenderedDisplayName"))
	return objectivec.Object{ID: rv}
}

// SiriVoiceGenderedDisplayName is an exported wrapper for the private method _siriVoiceGenderedDisplayName.
func (s SOVoiceObject) SiriVoiceGenderedDisplayName() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_siriVoiceGenderedDisplayName")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_siriVoiceGenderedDisplayName"}
		return nil, err
	}
	return s._siriVoiceGenderedDisplayName(), nil
}

// CanSiriVoiceGenderedDisplayName reports whether the receiver responds to the private selector _siriVoiceGenderedDisplayName.
func (s SOVoiceObject) CanSiriVoiceGenderedDisplayName() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_siriVoiceGenderedDisplayName"))
}
func (s SOVoiceObject) _siriVoiceGenderedDisplayNameRoot() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_siriVoiceGenderedDisplayNameRoot"))
	return objectivec.Object{ID: rv}
}

// SiriVoiceGenderedDisplayNameRoot is an exported wrapper for the private method _siriVoiceGenderedDisplayNameRoot.
func (s SOVoiceObject) SiriVoiceGenderedDisplayNameRoot() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_siriVoiceGenderedDisplayNameRoot")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_siriVoiceGenderedDisplayNameRoot"}
		return nil, err
	}
	return s._siriVoiceGenderedDisplayNameRoot(), nil
}

// CanSiriVoiceGenderedDisplayNameRoot reports whether the receiver responds to the private selector _siriVoiceGenderedDisplayNameRoot.
func (s SOVoiceObject) CanSiriVoiceGenderedDisplayNameRoot() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_siriVoiceGenderedDisplayNameRoot"))
}
func (s SOVoiceObject) _voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName:"), name)
	return objectivec.Object{ID: rv}
}

// VoiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName is an exported wrapper for the private method _voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName.
func (s SOVoiceObject) VoiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName(name objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName:"}
		return nil, err
	}
	return s._voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName(name), nil
}

// CanVoiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName reports whether the receiver responds to the private selector _voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName:.
func (s SOVoiceObject) CanVoiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_voiceNamesEntryFromSpeechSynthesisFrameworkForVoiceName:"))
}
func (s SOVoiceObject) Compare(compare objectivec.IObject) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("compare:"), compare)
	return rv
}
func (s SOVoiceObject) CountryDisplayString() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("countryDisplayString"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) CountryIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("countryIdentifier"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) DisplayName() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayName"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) DisplayNameRoot() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayNameRoot"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) DoesMatchSystemLocale() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("doesMatchSystemLocale"))
	return rv
}
func (s SOVoiceObject) Gender() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("gender"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) GenderDisplayString() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("genderDisplayString"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) Identifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("identifier"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) IsAppropriateForSystemLanguage() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isAppropriateForSystemLanguage"))
	return rv
}
func (s SOVoiceObject) IsNeuter() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isNeuter"))
	return rv
}
func (s SOVoiceObject) IsSiriVoice() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isSiriVoice"))
	return rv
}
func (s SOVoiceObject) LanguageDisplayString() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("languageDisplayString"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) LanguageIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("languageIdentifier"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) LocaleIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("localeIdentifier"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) MatchesSearchString(string_ objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("matchesSearchString:"), string_)
	return rv
}
func (s SOVoiceObject) RelativeDesirability() int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("relativeDesirability"))
	return rv
}
func (s SOVoiceObject) ShowsInFullListOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showsInFullListOnly"))
	return rv
}
func (s SOVoiceObject) SiriLocalizedColorName() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("siriLocalizedColorName"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) SiriVoiceDisplayNameFromIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("siriVoiceDisplayNameFromIdentifier"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) VoiceAttributes() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("voiceAttributes"))
	return objectivec.Object{ID: rv}
}
func (s SOVoiceObject) InitWithVoiceWithIdentifier(voice objectivec.IObject, identifier objectivec.IObject) SOVoiceObject {
	rv := objc.Send[SOVoiceObject](s.ID, objc.Sel("initWithVoice:identifier:"), voice, identifier)
	return rv
}
func (s SOVoiceObject) InitWithVoiceIdentifierWithIdentifier(identifier objectivec.IObject) SOVoiceObject {
	rv := objc.Send[SOVoiceObject](s.ID, objc.Sel("initWithVoiceIdentifier:"), identifier)
	return rv
}

func (_SOVoiceObjectClass SOVoiceObjectClass) AssetForVoiceID(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("assetForVoiceID:"), id)
	return objectivec.Object{ID: rv}
}
func (_SOVoiceObjectClass SOVoiceObjectClass) InvalidateAssetMaps() {
	objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("invalidateAssetMaps"))
}
func (_SOVoiceObjectClass SOVoiceObjectClass) IsSameLanguageFromLocaleIdentifierSecondLocaleIdentifier(identifier objectivec.IObject, identifier2 objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SOVoiceObjectClass.class), objc.Sel("isSameLanguageFromLocaleIdentifier:secondLocaleIdentifier:"), identifier, identifier2)
	return rv
}
func (_SOVoiceObjectClass SOVoiceObjectClass) IsSameLocaleIdentifierSecondLocaleIdentifier(identifier objectivec.IObject, identifier2 objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SOVoiceObjectClass.class), objc.Sel("isSameLocaleIdentifier:secondLocaleIdentifier:"), identifier, identifier2)
	return rv
}
func (_SOVoiceObjectClass SOVoiceObjectClass) NormalizedVoiceIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("normalizedVoiceIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
func (_SOVoiceObjectClass SOVoiceObjectClass) RebuildAssetMaps() {
	objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("rebuildAssetMaps"))
}
func (_SOVoiceObjectClass SOVoiceObjectClass) SetVisibleVoicesTable(table objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("setVisibleVoicesTable:"), table)
}
func (_SOVoiceObjectClass SOVoiceObjectClass) SystemLocaleIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("systemLocaleIdentifier"))
	return objectivec.Object{ID: rv}
}
func (_SOVoiceObjectClass SOVoiceObjectClass) VisibleVoicesForLocaleIdentifierAdditionalRequiredVoicesAllowAllVoices(identifier objectivec.IObject, voices objectivec.IObject, voices2 bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("visibleVoicesForLocaleIdentifier:additionalRequiredVoices:allowAllVoices:"), identifier, voices, voices2)
	return objectivec.Object{ID: rv}
}
func (_SOVoiceObjectClass SOVoiceObjectClass) VisibleVoicesTable() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOVoiceObjectClass.class), objc.Sel("visibleVoicesTable"))
	return objectivec.Object{ID: rv}
}

func (s SOVoiceObject) DownloadPercentComplete() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("downloadPercentComplete"))
	return rv
}
func (s SOVoiceObject) SetDownloadPercentComplete(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setDownloadPercentComplete:"), value)
}
func (s SOVoiceObject) DownloadStatus() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("downloadStatus"))
	return rv
}
func (s SOVoiceObject) SetDownloadStatus(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setDownloadStatus:"), value)
}
func (s SOVoiceObject) Visibility() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("visibility"))
	return rv
}
func (s SOVoiceObject) SetVisibility(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setVisibility:"), value)
}
func (s SOVoiceObject) Voice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("voice"))
	return rv
}
