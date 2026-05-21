// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechSynthesisProviderVoice] class.
var (
	_AVSpeechSynthesisProviderVoiceClass     AVSpeechSynthesisProviderVoiceClass
	_AVSpeechSynthesisProviderVoiceClassOnce sync.Once
)

func getAVSpeechSynthesisProviderVoiceClass() AVSpeechSynthesisProviderVoiceClass {
	_AVSpeechSynthesisProviderVoiceClassOnce.Do(func() {
		_AVSpeechSynthesisProviderVoiceClass = AVSpeechSynthesisProviderVoiceClass{class: objc.GetClass("AVSpeechSynthesisProviderVoice")}
	})
	return _AVSpeechSynthesisProviderVoiceClass
}

// GetAVSpeechSynthesisProviderVoiceClass returns the class object for AVSpeechSynthesisProviderVoice.
func GetAVSpeechSynthesisProviderVoiceClass() AVSpeechSynthesisProviderVoiceClass {
	return getAVSpeechSynthesisProviderVoiceClass()
}

type AVSpeechSynthesisProviderVoiceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechSynthesisProviderVoiceClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechSynthesisProviderVoiceClass) Alloc() AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVSpeechSynthesisProviderVoice.ExtraAttributes]
//   - [AVSpeechSynthesisProviderVoice.SetExtraAttributes]
//   - [AVSpeechSynthesisProviderVoice.FullBundleIdentifier]
//   - [AVSpeechSynthesisProviderVoice.GroupName]
//   - [AVSpeechSynthesisProviderVoice.IsFirstParty]
//   - [AVSpeechSynthesisProviderVoice.SetIsFirstParty]
//   - [AVSpeechSynthesisProviderVoice.IsPersonalVoice]
//   - [AVSpeechSynthesisProviderVoice.SetIsPersonalVoice]
//   - [AVSpeechSynthesisProviderVoice.ManufacturerName]
//   - [AVSpeechSynthesisProviderVoice.SetManufacturerName]
//   - [AVSpeechSynthesisProviderVoice.RawTTSIdentifier]
//   - [AVSpeechSynthesisProviderVoice.SupportedCharacterSet]
//   - [AVSpeechSynthesisProviderVoice.SynthesizerBundleIdentifier]
//   - [AVSpeechSynthesisProviderVoice.SetSynthesizerBundleIdentifier]
//   - [AVSpeechSynthesisProviderVoice.UniqueAudioDescSpeechSynthTuple]
//   - [AVSpeechSynthesisProviderVoice.UniqueAudioDescTriple]
//   - [AVSpeechSynthesisProviderVoice.Identifier]
//   - [AVSpeechSynthesisProviderVoice.SetIdentifier]
//   - [AVSpeechSynthesisProviderVoice.Name]
//   - [AVSpeechSynthesisProviderVoice.SetName]
//   - [AVSpeechSynthesisProviderVoice.PrimaryLanguages]
//   - [AVSpeechSynthesisProviderVoice.SetPrimaryLanguages]
//   - [AVSpeechSynthesisProviderVoice.SupportedLanguages]
//   - [AVSpeechSynthesisProviderVoice.SetSupportedLanguages]
type AVSpeechSynthesisProviderVoice struct {
	objectivec.Object
}

// AVSpeechSynthesisProviderVoiceFromID constructs a [AVSpeechSynthesisProviderVoice] from an objc.ID.
func AVSpeechSynthesisProviderVoiceFromID(id objc.ID) AVSpeechSynthesisProviderVoice {
	return AVSpeechSynthesisProviderVoice{objectivec.Object{ID: id}}
}

// Ensure AVSpeechSynthesisProviderVoice implements IAVSpeechSynthesisProviderVoice.
var _ IAVSpeechSynthesisProviderVoice = AVSpeechSynthesisProviderVoice{}

// An interface definition for the [AVSpeechSynthesisProviderVoice] class.
//
// # Methods
//
//   - [IAVSpeechSynthesisProviderVoice.ExtraAttributes]
//   - [IAVSpeechSynthesisProviderVoice.SetExtraAttributes]
//   - [IAVSpeechSynthesisProviderVoice.FullBundleIdentifier]
//   - [IAVSpeechSynthesisProviderVoice.GroupName]
//   - [IAVSpeechSynthesisProviderVoice.IsFirstParty]
//   - [IAVSpeechSynthesisProviderVoice.SetIsFirstParty]
//   - [IAVSpeechSynthesisProviderVoice.IsPersonalVoice]
//   - [IAVSpeechSynthesisProviderVoice.SetIsPersonalVoice]
//   - [IAVSpeechSynthesisProviderVoice.ManufacturerName]
//   - [IAVSpeechSynthesisProviderVoice.SetManufacturerName]
//   - [IAVSpeechSynthesisProviderVoice.RawTTSIdentifier]
//   - [IAVSpeechSynthesisProviderVoice.SupportedCharacterSet]
//   - [IAVSpeechSynthesisProviderVoice.SynthesizerBundleIdentifier]
//   - [IAVSpeechSynthesisProviderVoice.SetSynthesizerBundleIdentifier]
//   - [IAVSpeechSynthesisProviderVoice.UniqueAudioDescSpeechSynthTuple]
//   - [IAVSpeechSynthesisProviderVoice.UniqueAudioDescTriple]
//   - [IAVSpeechSynthesisProviderVoice.Identifier]
//   - [IAVSpeechSynthesisProviderVoice.SetIdentifier]
//   - [IAVSpeechSynthesisProviderVoice.Name]
//   - [IAVSpeechSynthesisProviderVoice.SetName]
//   - [IAVSpeechSynthesisProviderVoice.PrimaryLanguages]
//   - [IAVSpeechSynthesisProviderVoice.SetPrimaryLanguages]
//   - [IAVSpeechSynthesisProviderVoice.SupportedLanguages]
//   - [IAVSpeechSynthesisProviderVoice.SetSupportedLanguages]
type IAVSpeechSynthesisProviderVoice interface {
	objectivec.IObject

	// Topic: Methods

	ExtraAttributes() foundation.INSDictionary
	SetExtraAttributes(value foundation.INSDictionary)
	FullBundleIdentifier() objectivec.IObject
	GroupName() string
	IsFirstParty() bool
	SetIsFirstParty(value bool)
	IsPersonalVoice() bool
	SetIsPersonalVoice(value bool)
	ManufacturerName() string
	SetManufacturerName(value string)
	RawTTSIdentifier() objectivec.IObject
	SupportedCharacterSet() objectivec.IObject
	SynthesizerBundleIdentifier() string
	SetSynthesizerBundleIdentifier(value string)
	UniqueAudioDescSpeechSynthTuple() objectivec.IObject
	UniqueAudioDescTriple() objectivec.IObject
	Identifier() string
	SetIdentifier(value string)
	Name() string
	SetName(value string)
	PrimaryLanguages() foundation.INSArray
	SetPrimaryLanguages(value foundation.INSArray)
	SupportedLanguages() foundation.INSArray
	SetSupportedLanguages(value foundation.INSArray)
}

// Init initializes the instance.
func (a AVSpeechSynthesisProviderVoice) Init() AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVSpeechSynthesisProviderVoice) Autorelease() AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesisProviderVoice creates a new AVSpeechSynthesisProviderVoice instance.
func NewAVSpeechSynthesisProviderVoice() AVSpeechSynthesisProviderVoice {
	class := getAVSpeechSynthesisProviderVoiceClass()
	rv := objc.Send[AVSpeechSynthesisProviderVoice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVSpeechSynthesisProviderVoice) FullBundleIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fullBundleIdentifier"))
	return objectivec.Object{ID: rv}
}
func (a AVSpeechSynthesisProviderVoice) RawTTSIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("rawTTSIdentifier"))
	return objectivec.Object{ID: rv}
}
func (a AVSpeechSynthesisProviderVoice) SupportedCharacterSet() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("supportedCharacterSet"))
	return objectivec.Object{ID: rv}
}
func (a AVSpeechSynthesisProviderVoice) UniqueAudioDescSpeechSynthTuple() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("uniqueAudioDescSpeechSynthTuple"))
	return objectivec.Object{ID: rv}
}
func (a AVSpeechSynthesisProviderVoice) UniqueAudioDescTriple() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("uniqueAudioDescTriple"))
	return objectivec.Object{ID: rv}
}

func (_AVSpeechSynthesisProviderVoiceClass AVSpeechSynthesisProviderVoiceClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechSynthesisProviderVoiceClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
func (_AVSpeechSynthesisProviderVoiceClass AVSpeechSynthesisProviderVoiceClass) UpdateSpeechVoicesForClient(client objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisProviderVoiceClass.class), objc.Sel("updateSpeechVoicesForClient:"), client)
}

func (a AVSpeechSynthesisProviderVoice) ExtraAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("extraAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a AVSpeechSynthesisProviderVoice) SetExtraAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setExtraAttributes:"), value)
}
func (a AVSpeechSynthesisProviderVoice) GroupName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("groupName"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisProviderVoice) Identifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisProviderVoice) SetIdentifier(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setIdentifier:"), objc.String(value))
}
func (a AVSpeechSynthesisProviderVoice) IsFirstParty() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isFirstParty"))
	return rv
}
func (a AVSpeechSynthesisProviderVoice) SetIsFirstParty(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setIsFirstParty:"), value)
}
func (a AVSpeechSynthesisProviderVoice) IsPersonalVoice() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isPersonalVoice"))
	return rv
}
func (a AVSpeechSynthesisProviderVoice) SetIsPersonalVoice(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setIsPersonalVoice:"), value)
}
func (a AVSpeechSynthesisProviderVoice) ManufacturerName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("manufacturerName"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisProviderVoice) SetManufacturerName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setManufacturerName:"), objc.String(value))
}
func (a AVSpeechSynthesisProviderVoice) Name() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisProviderVoice) SetName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setName:"), objc.String(value))
}
func (a AVSpeechSynthesisProviderVoice) PrimaryLanguages() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("primaryLanguages"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a AVSpeechSynthesisProviderVoice) SetPrimaryLanguages(value foundation.INSArray) {
	objc.Send[struct{}](a.ID, objc.Sel("setPrimaryLanguages:"), value)
}
func (a AVSpeechSynthesisProviderVoice) SupportedLanguages() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("supportedLanguages"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a AVSpeechSynthesisProviderVoice) SetSupportedLanguages(value foundation.INSArray) {
	objc.Send[struct{}](a.ID, objc.Sel("setSupportedLanguages:"), value)
}
func (a AVSpeechSynthesisProviderVoice) SynthesizerBundleIdentifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("synthesizerBundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesisProviderVoice) SetSynthesizerBundleIdentifier(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setSynthesizerBundleIdentifier:"), objc.String(value))
}
