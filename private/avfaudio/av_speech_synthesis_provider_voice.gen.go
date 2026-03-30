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
//   - [AVSpeechSynthesisProviderVoice.InitWithCoder]
//   - [AVSpeechSynthesisProviderVoice.Identifier]
//   - [AVSpeechSynthesisProviderVoice.SetIdentifier]
//   - [AVSpeechSynthesisProviderVoice.Name]
//   - [AVSpeechSynthesisProviderVoice.SetName]
//   - [AVSpeechSynthesisProviderVoice.PrimaryLanguages]
//   - [AVSpeechSynthesisProviderVoice.SetPrimaryLanguages]
//   - [AVSpeechSynthesisProviderVoice.SupportedLanguages]
//   - [AVSpeechSynthesisProviderVoice.SetSupportedLanguages]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice
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
//   - [IAVSpeechSynthesisProviderVoice.InitWithCoder]
//   - [IAVSpeechSynthesisProviderVoice.Identifier]
//   - [IAVSpeechSynthesisProviderVoice.SetIdentifier]
//   - [IAVSpeechSynthesisProviderVoice.Name]
//   - [IAVSpeechSynthesisProviderVoice.SetName]
//   - [IAVSpeechSynthesisProviderVoice.PrimaryLanguages]
//   - [IAVSpeechSynthesisProviderVoice.SetPrimaryLanguages]
//   - [IAVSpeechSynthesisProviderVoice.SupportedLanguages]
//   - [IAVSpeechSynthesisProviderVoice.SetSupportedLanguages]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice
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
	InitWithCoder(coder foundation.INSCoder) AVSpeechSynthesisProviderVoice
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
func (s AVSpeechSynthesisProviderVoice) Init() AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechSynthesisProviderVoice) Autorelease() AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesisProviderVoice creates a new AVSpeechSynthesisProviderVoice instance.
func NewAVSpeechSynthesisProviderVoice() AVSpeechSynthesisProviderVoice {
	class := getAVSpeechSynthesisProviderVoiceClass()
	rv := objc.Send[AVSpeechSynthesisProviderVoice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/initWithCoder:
func NewSpeechSynthesisProviderVoiceWithCoder(coder objectivec.IObject) AVSpeechSynthesisProviderVoice {
	instance := getAVSpeechSynthesisProviderVoiceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AVSpeechSynthesisProviderVoiceFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/fullBundleIdentifier
func (s AVSpeechSynthesisProviderVoice) FullBundleIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fullBundleIdentifier"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/rawTTSIdentifier
func (s AVSpeechSynthesisProviderVoice) RawTTSIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("rawTTSIdentifier"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/supportedCharacterSet
func (s AVSpeechSynthesisProviderVoice) SupportedCharacterSet() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedCharacterSet"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/uniqueAudioDescSpeechSynthTuple
func (s AVSpeechSynthesisProviderVoice) UniqueAudioDescSpeechSynthTuple() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("uniqueAudioDescSpeechSynthTuple"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/uniqueAudioDescTriple
func (s AVSpeechSynthesisProviderVoice) UniqueAudioDescTriple() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("uniqueAudioDescTriple"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/initWithCoder:
func (s AVSpeechSynthesisProviderVoice) InitWithCoder(coder foundation.INSCoder) AVSpeechSynthesisProviderVoice {
	rv := objc.Send[AVSpeechSynthesisProviderVoice](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/supportsSecureCoding
func (_AVSpeechSynthesisProviderVoiceClass AVSpeechSynthesisProviderVoiceClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechSynthesisProviderVoiceClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/updateSpeechVoicesForClient:
func (_AVSpeechSynthesisProviderVoiceClass AVSpeechSynthesisProviderVoiceClass) UpdateSpeechVoicesForClient(client objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_AVSpeechSynthesisProviderVoiceClass.class), objc.Sel("updateSpeechVoicesForClient:"), client)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/extraAttributes
func (s AVSpeechSynthesisProviderVoice) ExtraAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("extraAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s AVSpeechSynthesisProviderVoice) SetExtraAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setExtraAttributes:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/groupName
func (s AVSpeechSynthesisProviderVoice) GroupName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("groupName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/identifier
func (s AVSpeechSynthesisProviderVoice) Identifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisProviderVoice) SetIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/isFirstParty
func (s AVSpeechSynthesisProviderVoice) IsFirstParty() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isFirstParty"))
	return rv
}
func (s AVSpeechSynthesisProviderVoice) SetIsFirstParty(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIsFirstParty:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/isPersonalVoice
func (s AVSpeechSynthesisProviderVoice) IsPersonalVoice() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isPersonalVoice"))
	return rv
}
func (s AVSpeechSynthesisProviderVoice) SetIsPersonalVoice(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIsPersonalVoice:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/manufacturerName
func (s AVSpeechSynthesisProviderVoice) ManufacturerName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("manufacturerName"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisProviderVoice) SetManufacturerName(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setManufacturerName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/name
func (s AVSpeechSynthesisProviderVoice) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisProviderVoice) SetName(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/primaryLanguages
func (s AVSpeechSynthesisProviderVoice) PrimaryLanguages() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("primaryLanguages"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s AVSpeechSynthesisProviderVoice) SetPrimaryLanguages(value foundation.INSArray) {
	objc.Send[struct{}](s.ID, objc.Sel("setPrimaryLanguages:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/supportedLanguages
func (s AVSpeechSynthesisProviderVoice) SupportedLanguages() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportedLanguages"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s AVSpeechSynthesisProviderVoice) SetSupportedLanguages(value foundation.INSArray) {
	objc.Send[struct{}](s.ID, objc.Sel("setSupportedLanguages:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisProviderVoice/synthesizerBundleIdentifier
func (s AVSpeechSynthesisProviderVoice) SynthesizerBundleIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("synthesizerBundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisProviderVoice) SetSynthesizerBundleIdentifier(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setSynthesizerBundleIdentifier:"), objc.String(value))
}
