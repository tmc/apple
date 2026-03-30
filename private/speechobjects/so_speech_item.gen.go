// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOSpeechItem] class.
var (
	_SOSpeechItemClass     SOSpeechItemClass
	_SOSpeechItemClassOnce sync.Once
)

func getSOSpeechItemClass() SOSpeechItemClass {
	_SOSpeechItemClassOnce.Do(func() {
		_SOSpeechItemClass = SOSpeechItemClass{class: objc.GetClass("SOSpeechItem")}
	})
	return _SOSpeechItemClass
}

// GetSOSpeechItemClass returns the class object for SOSpeechItem.
func GetSOSpeechItemClass() SOSpeechItemClass {
	return getSOSpeechItemClass()
}

type SOSpeechItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOSpeechItemClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOSpeechItemClass) Alloc() SOSpeechItem {
	rv := objc.Send[SOSpeechItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOSpeechItem._conversionLocale]
//   - [SOSpeechItem.BundleIdentifier]
//   - [SOSpeechItem.ByteSize]
//   - [SOSpeechItem.Compare]
//   - [SOSpeechItem.CompareCountryTitle]
//   - [SOSpeechItem.CompareCountryWithDialectTitle]
//   - [SOSpeechItem.CompareDisplayTitle]
//   - [SOSpeechItem.CompareLanguageTitle]
//   - [SOSpeechItem.CountryDisplayString]
//   - [SOSpeechItem.CountryIdentifier]
//   - [SOSpeechItem.CountryWithDialectDisplayString]
//   - [SOSpeechItem.DisplayTitle]
//   - [SOSpeechItem.DisplayedSize]
//   - [SOSpeechItem.DoesMatchSystemLocale]
//   - [SOSpeechItem.DownloadFullSize]
//   - [SOSpeechItem.SetDownloadFullSize]
//   - [SOSpeechItem.DownloadPercentComplete]
//   - [SOSpeechItem.SetDownloadPercentComplete]
//   - [SOSpeechItem.DownloadStatus]
//   - [SOSpeechItem.SetDownloadStatus]
//   - [SOSpeechItem.EngineIdentifier]
//   - [SOSpeechItem.FallbackLocaleIdentifier]
//   - [SOSpeechItem.SetFallbackLocaleIdentifier]
//   - [SOSpeechItem.FullSizeBundleIdentifier]
//   - [SOSpeechItem.FullSizeByteSize]
//   - [SOSpeechItem.FullSizeTagName]
//   - [SOSpeechItem.IsAppropriateForSystemLangauge]
//   - [SOSpeechItem.LanguageDisplayString]
//   - [SOSpeechItem.LanguageIdentifier]
//   - [SOSpeechItem.LanguageWithDialectDisplayString]
//   - [SOSpeechItem.LocaleIdentifier]
//   - [SOSpeechItem.SetLocaleIdentifier]
//   - [SOSpeechItem.MatchesSearchString]
//   - [SOSpeechItem.OfflineDictationOnly]
//   - [SOSpeechItem.SetOfflineDictationOnly]
//   - [SOSpeechItem.ShowCountryInDisplayTitle]
//   - [SOSpeechItem.SetShowCountryInDisplayTitle]
//   - [SOSpeechItem.TagName]
//   - [SOSpeechItem.Variant]
//   - [SOSpeechItem.Visibility]
//   - [SOSpeechItem.SetVisibility]
//   - [SOSpeechItem.InitWithDownloadableBundleIdentifierProperties]
//   - [SOSpeechItem.Version]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem
type SOSpeechItem struct {
	objectivec.Object
}

// SOSpeechItemFromID constructs a [SOSpeechItem] from an objc.ID.
func SOSpeechItemFromID(id objc.ID) SOSpeechItem {
	return SOSpeechItem{objectivec.Object{ID: id}}
}

// Ensure SOSpeechItem implements ISOSpeechItem.
var _ ISOSpeechItem = SOSpeechItem{}

// An interface definition for the [SOSpeechItem] class.
//
// # Methods
//
//   - [ISOSpeechItem._conversionLocale]
//   - [ISOSpeechItem.BundleIdentifier]
//   - [ISOSpeechItem.ByteSize]
//   - [ISOSpeechItem.Compare]
//   - [ISOSpeechItem.CompareCountryTitle]
//   - [ISOSpeechItem.CompareCountryWithDialectTitle]
//   - [ISOSpeechItem.CompareDisplayTitle]
//   - [ISOSpeechItem.CompareLanguageTitle]
//   - [ISOSpeechItem.CountryDisplayString]
//   - [ISOSpeechItem.CountryIdentifier]
//   - [ISOSpeechItem.CountryWithDialectDisplayString]
//   - [ISOSpeechItem.DisplayTitle]
//   - [ISOSpeechItem.DisplayedSize]
//   - [ISOSpeechItem.DoesMatchSystemLocale]
//   - [ISOSpeechItem.DownloadFullSize]
//   - [ISOSpeechItem.SetDownloadFullSize]
//   - [ISOSpeechItem.DownloadPercentComplete]
//   - [ISOSpeechItem.SetDownloadPercentComplete]
//   - [ISOSpeechItem.DownloadStatus]
//   - [ISOSpeechItem.SetDownloadStatus]
//   - [ISOSpeechItem.EngineIdentifier]
//   - [ISOSpeechItem.FallbackLocaleIdentifier]
//   - [ISOSpeechItem.SetFallbackLocaleIdentifier]
//   - [ISOSpeechItem.FullSizeBundleIdentifier]
//   - [ISOSpeechItem.FullSizeByteSize]
//   - [ISOSpeechItem.FullSizeTagName]
//   - [ISOSpeechItem.IsAppropriateForSystemLangauge]
//   - [ISOSpeechItem.LanguageDisplayString]
//   - [ISOSpeechItem.LanguageIdentifier]
//   - [ISOSpeechItem.LanguageWithDialectDisplayString]
//   - [ISOSpeechItem.LocaleIdentifier]
//   - [ISOSpeechItem.SetLocaleIdentifier]
//   - [ISOSpeechItem.MatchesSearchString]
//   - [ISOSpeechItem.OfflineDictationOnly]
//   - [ISOSpeechItem.SetOfflineDictationOnly]
//   - [ISOSpeechItem.ShowCountryInDisplayTitle]
//   - [ISOSpeechItem.SetShowCountryInDisplayTitle]
//   - [ISOSpeechItem.TagName]
//   - [ISOSpeechItem.Variant]
//   - [ISOSpeechItem.Visibility]
//   - [ISOSpeechItem.SetVisibility]
//   - [ISOSpeechItem.InitWithDownloadableBundleIdentifierProperties]
//   - [ISOSpeechItem.Version]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem
type ISOSpeechItem interface {
	objectivec.IObject

	// Topic: Methods

	_conversionLocale() objectivec.IObject
	BundleIdentifier() string
	ByteSize() uint64
	Compare(compare objectivec.IObject) int64
	CompareCountryTitle(title objectivec.IObject) int64
	CompareCountryWithDialectTitle(title objectivec.IObject) int64
	CompareDisplayTitle(title objectivec.IObject) int64
	CompareLanguageTitle(title objectivec.IObject) int64
	CountryDisplayString() objectivec.IObject
	CountryIdentifier() objectivec.IObject
	CountryWithDialectDisplayString() objectivec.IObject
	DisplayTitle() objectivec.IObject
	DisplayedSize() objectivec.IObject
	DoesMatchSystemLocale() bool
	DownloadFullSize() bool
	SetDownloadFullSize(value bool)
	DownloadPercentComplete() float64
	SetDownloadPercentComplete(value float64)
	DownloadStatus() uint64
	SetDownloadStatus(value uint64)
	EngineIdentifier() string
	FallbackLocaleIdentifier() string
	SetFallbackLocaleIdentifier(value string)
	FullSizeBundleIdentifier() string
	FullSizeByteSize() uint64
	FullSizeTagName() string
	IsAppropriateForSystemLangauge() bool
	LanguageDisplayString() objectivec.IObject
	LanguageIdentifier() objectivec.IObject
	LanguageWithDialectDisplayString() objectivec.IObject
	LocaleIdentifier() string
	SetLocaleIdentifier(value string)
	MatchesSearchString(string_ objectivec.IObject) bool
	OfflineDictationOnly() bool
	SetOfflineDictationOnly(value bool)
	ShowCountryInDisplayTitle() bool
	SetShowCountryInDisplayTitle(value bool)
	TagName() string
	Variant() string
	Visibility() uint64
	SetVisibility(value uint64)
	InitWithDownloadableBundleIdentifierProperties(identifier objectivec.IObject, properties objectivec.IObject) SOSpeechItem
	Version() string
}

// Init initializes the instance.
func (s SOSpeechItem) Init() SOSpeechItem {
	rv := objc.Send[SOSpeechItem](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOSpeechItem) Autorelease() SOSpeechItem {
	rv := objc.Send[SOSpeechItem](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOSpeechItem creates a new SOSpeechItem instance.
func NewSOSpeechItem() SOSpeechItem {
	class := getSOSpeechItemClass()
	rv := objc.Send[SOSpeechItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/initWithDownloadableBundleIdentifier:properties:
func NewSOSpeechItemWithDownloadableBundleIdentifierProperties(identifier objectivec.IObject, properties objectivec.IObject) SOSpeechItem {
	instance := getSOSpeechItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDownloadableBundleIdentifier:properties:"), identifier, properties)
	return SOSpeechItemFromID(rv)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/_conversionLocale
func (s SOSpeechItem) _conversionLocale() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_conversionLocale"))
	return objectivec.Object{ID: rv}
}

// ConversionLocale is an exported wrapper for the private method _conversionLocale.
func (s SOSpeechItem) ConversionLocale() objectivec.IObject {
	return s._conversionLocale()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/compare:
func (s SOSpeechItem) Compare(compare objectivec.IObject) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("compare:"), compare)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/compareCountryTitle:
func (s SOSpeechItem) CompareCountryTitle(title objectivec.IObject) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("compareCountryTitle:"), title)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/compareCountryWithDialectTitle:
func (s SOSpeechItem) CompareCountryWithDialectTitle(title objectivec.IObject) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("compareCountryWithDialectTitle:"), title)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/compareDisplayTitle:
func (s SOSpeechItem) CompareDisplayTitle(title objectivec.IObject) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("compareDisplayTitle:"), title)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/compareLanguageTitle:
func (s SOSpeechItem) CompareLanguageTitle(title objectivec.IObject) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("compareLanguageTitle:"), title)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/countryDisplayString
func (s SOSpeechItem) CountryDisplayString() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("countryDisplayString"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/countryIdentifier
func (s SOSpeechItem) CountryIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("countryIdentifier"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/countryWithDialectDisplayString
func (s SOSpeechItem) CountryWithDialectDisplayString() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("countryWithDialectDisplayString"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/displayTitle
func (s SOSpeechItem) DisplayTitle() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayTitle"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/displayedSize
func (s SOSpeechItem) DisplayedSize() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayedSize"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/doesMatchSystemLocale
func (s SOSpeechItem) DoesMatchSystemLocale() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("doesMatchSystemLocale"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/isAppropriateForSystemLangauge
func (s SOSpeechItem) IsAppropriateForSystemLangauge() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isAppropriateForSystemLangauge"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/languageDisplayString
func (s SOSpeechItem) LanguageDisplayString() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("languageDisplayString"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/languageIdentifier
func (s SOSpeechItem) LanguageIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("languageIdentifier"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/languageWithDialectDisplayString
func (s SOSpeechItem) LanguageWithDialectDisplayString() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("languageWithDialectDisplayString"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/matchesSearchString:
func (s SOSpeechItem) MatchesSearchString(string_ objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("matchesSearchString:"), string_)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/initWithDownloadableBundleIdentifier:properties:
func (s SOSpeechItem) InitWithDownloadableBundleIdentifierProperties(identifier objectivec.IObject, properties objectivec.IObject) SOSpeechItem {
	rv := objc.Send[SOSpeechItem](s.ID, objc.Sel("initWithDownloadableBundleIdentifier:properties:"), identifier, properties)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/componentsFromLocaleIdentifier:
func (_SOSpeechItemClass SOSpeechItemClass) ComponentsFromLocaleIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOSpeechItemClass.class), objc.Sel("componentsFromLocaleIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/displayStringForByteSize:
func (_SOSpeechItemClass SOSpeechItemClass) DisplayStringForByteSize(size uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOSpeechItemClass.class), objc.Sel("displayStringForByteSize:"), size)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/displayStringForTimeRemaining:
func (_SOSpeechItemClass SOSpeechItemClass) DisplayStringForTimeRemaining(remaining float64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOSpeechItemClass.class), objc.Sel("displayStringForTimeRemaining:"), remaining)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/installationDictionaryCache
func (_SOSpeechItemClass SOSpeechItemClass) InstallationDictionaryCache() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOSpeechItemClass.class), objc.Sel("installationDictionaryCache"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/isLocaleIdentifier:containedInLocaleIdentifiers:
func (_SOSpeechItemClass SOSpeechItemClass) IsLocaleIdentifierContainedInLocaleIdentifiers(identifier objectivec.IObject, identifiers objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SOSpeechItemClass.class), objc.Sel("isLocaleIdentifier:containedInLocaleIdentifiers:"), identifier, identifiers)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/isSameLangaugeFromLocaleIdentifier:secondLocaleIdentifier:
func (_SOSpeechItemClass SOSpeechItemClass) IsSameLangaugeFromLocaleIdentifierSecondLocaleIdentifier(identifier objectivec.IObject, identifier2 objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SOSpeechItemClass.class), objc.Sel("isSameLangaugeFromLocaleIdentifier:secondLocaleIdentifier:"), identifier, identifier2)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/isSameLocaleIdentifier:secondLocaleIdentifier:
func (_SOSpeechItemClass SOSpeechItemClass) IsSameLocaleIdentifierSecondLocaleIdentifier(identifier objectivec.IObject, identifier2 objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SOSpeechItemClass.class), objc.Sel("isSameLocaleIdentifier:secondLocaleIdentifier:"), identifier, identifier2)
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/normalizedLocaleIdentifier:
func (_SOSpeechItemClass SOSpeechItemClass) NormalizedLocaleIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOSpeechItemClass.class), objc.Sel("normalizedLocaleIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/systemLocaleIdentifier
func (_SOSpeechItemClass SOSpeechItemClass) SystemLocaleIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOSpeechItemClass.class), objc.Sel("systemLocaleIdentifier"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/visibleObjectsFromItems:forLocaleIdentifier:additionalRequiredItems:
func (_SOSpeechItemClass SOSpeechItemClass) VisibleObjectsFromItemsForLocaleIdentifierAdditionalRequiredItems(items objectivec.IObject, identifier objectivec.IObject, items2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SOSpeechItemClass.class), objc.Sel("visibleObjectsFromItems:forLocaleIdentifier:additionalRequiredItems:"), items, identifier, items2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/bundleIdentifier
func (s SOSpeechItem) BundleIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/byteSize
func (s SOSpeechItem) ByteSize() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("byteSize"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/downloadFullSize
func (s SOSpeechItem) DownloadFullSize() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("downloadFullSize"))
	return rv
}
func (s SOSpeechItem) SetDownloadFullSize(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setDownloadFullSize:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/downloadPercentComplete
func (s SOSpeechItem) DownloadPercentComplete() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("downloadPercentComplete"))
	return rv
}
func (s SOSpeechItem) SetDownloadPercentComplete(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setDownloadPercentComplete:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/downloadStatus
func (s SOSpeechItem) DownloadStatus() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("downloadStatus"))
	return rv
}
func (s SOSpeechItem) SetDownloadStatus(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setDownloadStatus:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/engineIdentifier
func (s SOSpeechItem) EngineIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("engineIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/fallbackLocaleIdentifier
func (s SOSpeechItem) FallbackLocaleIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fallbackLocaleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOSpeechItem) SetFallbackLocaleIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setFallbackLocaleIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/fullSizeBundleIdentifier
func (s SOSpeechItem) FullSizeBundleIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fullSizeBundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/fullSizeByteSize
func (s SOSpeechItem) FullSizeByteSize() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("fullSizeByteSize"))
	return rv
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/fullSizeTagName
func (s SOSpeechItem) FullSizeTagName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fullSizeTagName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/localeIdentifier
func (s SOSpeechItem) LocaleIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("localeIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s SOSpeechItem) SetLocaleIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setLocaleIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/offlineDictationOnly
func (s SOSpeechItem) OfflineDictationOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("offlineDictationOnly"))
	return rv
}
func (s SOSpeechItem) SetOfflineDictationOnly(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setOfflineDictationOnly:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/showCountryInDisplayTitle
func (s SOSpeechItem) ShowCountryInDisplayTitle() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showCountryInDisplayTitle"))
	return rv
}
func (s SOSpeechItem) SetShowCountryInDisplayTitle(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShowCountryInDisplayTitle:"), value)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/tagName
func (s SOSpeechItem) TagName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("tagName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/variant
func (s SOSpeechItem) Variant() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("variant"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/version
func (s SOSpeechItem) Version() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("version"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOSpeechItem/visibility
func (s SOSpeechItem) Visibility() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("visibility"))
	return rv
}
func (s SOSpeechItem) SetVisibility(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setVisibility:"), value)
}
