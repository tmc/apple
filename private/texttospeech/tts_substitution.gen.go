// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSSubstitution] class.
var (
	_TTSSubstitutionClass     TTSSubstitutionClass
	_TTSSubstitutionClassOnce sync.Once
)

func getTTSSubstitutionClass() TTSSubstitutionClass {
	_TTSSubstitutionClassOnce.Do(func() {
		_TTSSubstitutionClass = TTSSubstitutionClass{class: objc.GetClass("TTSSubstitution")}
	})
	return _TTSSubstitutionClass
}

// GetTTSSubstitutionClass returns the class object for TTSSubstitution.
func GetTTSSubstitutionClass() TTSSubstitutionClass {
	return getTTSSubstitutionClass()
}

type TTSSubstitutionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSubstitutionClass) Alloc() TTSSubstitution {
	rv := objc.Send[TTSSubstitution](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSSubstitution.Active]
//   - [TTSSubstitution.SetActive]
//   - [TTSSubstitution.AppliesToAllApps]
//   - [TTSSubstitution.SetAppliesToAllApps]
//   - [TTSSubstitution.BundleIdentifiers]
//   - [TTSSubstitution.SetBundleIdentifiers]
//   - [TTSSubstitution.EncodeWithCoder]
//   - [TTSSubstitution.IgnoreCase]
//   - [TTSSubstitution.SetIgnoreCase]
//   - [TTSSubstitution.IsReplacementTextAllPunctuation]
//   - [TTSSubstitution.IsReplacementTextSurroundedByPunctuation]
//   - [TTSSubstitution.IsUserSubstitution]
//   - [TTSSubstitution.SetIsUserSubstitution]
//   - [TTSSubstitution.Languages]
//   - [TTSSubstitution.SetLanguages]
//   - [TTSSubstitution.OriginalString]
//   - [TTSSubstitution.SetOriginalString]
//   - [TTSSubstitution.Phonemes]
//   - [TTSSubstitution.SetPhonemes]
//   - [TTSSubstitution.ReplacementRange]
//   - [TTSSubstitution.SetReplacementRange]
//   - [TTSSubstitution.ReplacementString]
//   - [TTSSubstitution.SetReplacementString]
//   - [TTSSubstitution.Uuid]
//   - [TTSSubstitution.SetUuid]
//   - [TTSSubstitution.VoiceIds]
//   - [TTSSubstitution.SetVoiceIds]
//   - [TTSSubstitution.InitWithCoder]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution
type TTSSubstitution struct {
	objectivec.Object
}

// TTSSubstitutionFromID constructs a [TTSSubstitution] from an objc.ID.
func TTSSubstitutionFromID(id objc.ID) TTSSubstitution {
	return TTSSubstitution{objectivec.Object{ID: id}}
}
// Ensure TTSSubstitution implements ITTSSubstitution.
var _ ITTSSubstitution = TTSSubstitution{}

// An interface definition for the [TTSSubstitution] class.
//
// # Methods
//
//   - [ITTSSubstitution.Active]
//   - [ITTSSubstitution.SetActive]
//   - [ITTSSubstitution.AppliesToAllApps]
//   - [ITTSSubstitution.SetAppliesToAllApps]
//   - [ITTSSubstitution.BundleIdentifiers]
//   - [ITTSSubstitution.SetBundleIdentifiers]
//   - [ITTSSubstitution.EncodeWithCoder]
//   - [ITTSSubstitution.IgnoreCase]
//   - [ITTSSubstitution.SetIgnoreCase]
//   - [ITTSSubstitution.IsReplacementTextAllPunctuation]
//   - [ITTSSubstitution.IsReplacementTextSurroundedByPunctuation]
//   - [ITTSSubstitution.IsUserSubstitution]
//   - [ITTSSubstitution.SetIsUserSubstitution]
//   - [ITTSSubstitution.Languages]
//   - [ITTSSubstitution.SetLanguages]
//   - [ITTSSubstitution.OriginalString]
//   - [ITTSSubstitution.SetOriginalString]
//   - [ITTSSubstitution.Phonemes]
//   - [ITTSSubstitution.SetPhonemes]
//   - [ITTSSubstitution.ReplacementRange]
//   - [ITTSSubstitution.SetReplacementRange]
//   - [ITTSSubstitution.ReplacementString]
//   - [ITTSSubstitution.SetReplacementString]
//   - [ITTSSubstitution.Uuid]
//   - [ITTSSubstitution.SetUuid]
//   - [ITTSSubstitution.VoiceIds]
//   - [ITTSSubstitution.SetVoiceIds]
//   - [ITTSSubstitution.InitWithCoder]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution
type ITTSSubstitution interface {
	objectivec.IObject

	// Topic: Methods

	Active() bool
	SetActive(value bool)
	AppliesToAllApps() bool
	SetAppliesToAllApps(value bool)
	BundleIdentifiers() foundation.INSSet
	SetBundleIdentifiers(value foundation.INSSet)
	EncodeWithCoder(coder foundation.INSCoder)
	IgnoreCase() bool
	SetIgnoreCase(value bool)
	IsReplacementTextAllPunctuation() bool
	IsReplacementTextSurroundedByPunctuation() bool
	IsUserSubstitution() bool
	SetIsUserSubstitution(value bool)
	Languages() foundation.INSSet
	SetLanguages(value foundation.INSSet)
	OriginalString() string
	SetOriginalString(value string)
	Phonemes() string
	SetPhonemes(value string)
	ReplacementRange() foundation.NSRange
	SetReplacementRange(value foundation.NSRange)
	ReplacementString() string
	SetReplacementString(value string)
	Uuid() foundation.NSUUID
	SetUuid(value foundation.NSUUID)
	VoiceIds() foundation.INSSet
	SetVoiceIds(value foundation.INSSet)
	InitWithCoder(coder foundation.INSCoder) TTSSubstitution
}

// Init initializes the instance.
func (t TTSSubstitution) Init() TTSSubstitution {
	rv := objc.Send[TTSSubstitution](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSubstitution) Autorelease() TTSSubstitution {
	rv := objc.Send[TTSSubstitution](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSubstitution creates a new TTSSubstitution instance.
func NewTTSSubstitution() TTSSubstitution {
	class := getTTSSubstitutionClass()
	rv := objc.Send[TTSSubstitution](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/initWithCoder:
func NewTTSSubstitutionWithCoder(coder objectivec.IObject) TTSSubstitution {
	instance := getTTSSubstitutionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return TTSSubstitutionFromID(rv)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/encodeWithCoder:
func (t TTSSubstitution) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/initWithCoder:
func (t TTSSubstitution) InitWithCoder(coder foundation.INSCoder) TTSSubstitution {
	rv := objc.Send[TTSSubstitution](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/supportsSecureCoding
func (_TTSSubstitutionClass TTSSubstitutionClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_TTSSubstitutionClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/active
func (t TTSSubstitution) Active() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("active"))
	return rv
}
func (t TTSSubstitution) SetActive(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setActive:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/appliesToAllApps
func (t TTSSubstitution) AppliesToAllApps() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("appliesToAllApps"))
	return rv
}
func (t TTSSubstitution) SetAppliesToAllApps(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAppliesToAllApps:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/bundleIdentifiers
func (t TTSSubstitution) BundleIdentifiers() foundation.INSSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("bundleIdentifiers"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (t TTSSubstitution) SetBundleIdentifiers(value foundation.INSSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setBundleIdentifiers:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/ignoreCase
func (t TTSSubstitution) IgnoreCase() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("ignoreCase"))
	return rv
}
func (t TTSSubstitution) SetIgnoreCase(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIgnoreCase:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/isReplacementTextAllPunctuation
func (t TTSSubstitution) IsReplacementTextAllPunctuation() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isReplacementTextAllPunctuation"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/isReplacementTextSurroundedByPunctuation
func (t TTSSubstitution) IsReplacementTextSurroundedByPunctuation() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isReplacementTextSurroundedByPunctuation"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/isUserSubstitution
func (t TTSSubstitution) IsUserSubstitution() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isUserSubstitution"))
	return rv
}
func (t TTSSubstitution) SetIsUserSubstitution(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIsUserSubstitution:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/languages
func (t TTSSubstitution) Languages() foundation.INSSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("languages"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (t TTSSubstitution) SetLanguages(value foundation.INSSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setLanguages:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/originalString
func (t TTSSubstitution) OriginalString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("originalString"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSubstitution) SetOriginalString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setOriginalString:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/phonemes
func (t TTSSubstitution) Phonemes() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("phonemes"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSubstitution) SetPhonemes(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPhonemes:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/replacementRange
func (t TTSSubstitution) ReplacementRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("replacementRange"))
	return foundation.NSRange(rv)
}
func (t TTSSubstitution) SetReplacementRange(value foundation.NSRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setReplacementRange:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/replacementString
func (t TTSSubstitution) ReplacementString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("replacementString"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSubstitution) SetReplacementString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setReplacementString:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/uuid
func (t TTSSubstitution) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (t TTSSubstitution) SetUuid(value foundation.NSUUID) {
	objc.Send[struct{}](t.ID, objc.Sel("setUuid:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSubstitution/voiceIds
func (t TTSSubstitution) VoiceIds() foundation.INSSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceIds"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (t TTSSubstitution) SetVoiceIds(value foundation.INSSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoiceIds:"), value)
}

