// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAXResourceMigrationUtilities] class.
var (
	_TTSAXResourceMigrationUtilitiesClass     TTSAXResourceMigrationUtilitiesClass
	_TTSAXResourceMigrationUtilitiesClassOnce sync.Once
)

func getTTSAXResourceMigrationUtilitiesClass() TTSAXResourceMigrationUtilitiesClass {
	_TTSAXResourceMigrationUtilitiesClassOnce.Do(func() {
		_TTSAXResourceMigrationUtilitiesClass = TTSAXResourceMigrationUtilitiesClass{class: objc.GetClass("TTSAXResourceMigrationUtilities")}
	})
	return _TTSAXResourceMigrationUtilitiesClass
}

// GetTTSAXResourceMigrationUtilitiesClass returns the class object for TTSAXResourceMigrationUtilities.
func GetTTSAXResourceMigrationUtilitiesClass() TTSAXResourceMigrationUtilitiesClass {
	return getTTSAXResourceMigrationUtilitiesClass()
}

type TTSAXResourceMigrationUtilitiesClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAXResourceMigrationUtilitiesClass) Alloc() TTSAXResourceMigrationUtilities {
	rv := objc.Send[TTSAXResourceMigrationUtilities](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSAXResourceMigrationUtilities._ttsAssetForSpec]
//   - [TTSAXResourceMigrationUtilities.AttributesForLegacyVoiceSpec]
//   - [TTSAXResourceMigrationUtilities.GetIdentifierForLegacyVoiceSpec]
//   - [TTSAXResourceMigrationUtilities.TtsIdentifierForLegacyMacIdentifier]
//   - [TTSAXResourceMigrationUtilities.UpdatedIdentifierForLegacyIdentifierWithLanguageCode]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceMigrationUtilities
type TTSAXResourceMigrationUtilities struct {
	objectivec.Object
}

// TTSAXResourceMigrationUtilitiesFromID constructs a [TTSAXResourceMigrationUtilities] from an objc.ID.
func TTSAXResourceMigrationUtilitiesFromID(id objc.ID) TTSAXResourceMigrationUtilities {
	return TTSAXResourceMigrationUtilities{objectivec.Object{ID: id}}
}
// Ensure TTSAXResourceMigrationUtilities implements ITTSAXResourceMigrationUtilities.
var _ ITTSAXResourceMigrationUtilities = TTSAXResourceMigrationUtilities{}

// An interface definition for the [TTSAXResourceMigrationUtilities] class.
//
// # Methods
//
//   - [ITTSAXResourceMigrationUtilities._ttsAssetForSpec]
//   - [ITTSAXResourceMigrationUtilities.AttributesForLegacyVoiceSpec]
//   - [ITTSAXResourceMigrationUtilities.GetIdentifierForLegacyVoiceSpec]
//   - [ITTSAXResourceMigrationUtilities.TtsIdentifierForLegacyMacIdentifier]
//   - [ITTSAXResourceMigrationUtilities.UpdatedIdentifierForLegacyIdentifierWithLanguageCode]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceMigrationUtilities
type ITTSAXResourceMigrationUtilities interface {
	objectivec.IObject

	// Topic: Methods

	_ttsAssetForSpec(spec objectivec.IObject) objectivec.IObject
	AttributesForLegacyVoiceSpec(spec objectivec.IObject) objectivec.IObject
	GetIdentifierForLegacyVoiceSpec(spec objectivec.IObject) objectivec.IObject
	TtsIdentifierForLegacyMacIdentifier(identifier objectivec.IObject) objectivec.IObject
	UpdatedIdentifierForLegacyIdentifierWithLanguageCode(identifier objectivec.IObject, code objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (t TTSAXResourceMigrationUtilities) Init() TTSAXResourceMigrationUtilities {
	rv := objc.Send[TTSAXResourceMigrationUtilities](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAXResourceMigrationUtilities) Autorelease() TTSAXResourceMigrationUtilities {
	rv := objc.Send[TTSAXResourceMigrationUtilities](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAXResourceMigrationUtilities creates a new TTSAXResourceMigrationUtilities instance.
func NewTTSAXResourceMigrationUtilities() TTSAXResourceMigrationUtilities {
	class := getTTSAXResourceMigrationUtilitiesClass()
	rv := objc.Send[TTSAXResourceMigrationUtilities](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// spec is a [applicationservices.VoiceSpec].
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceMigrationUtilities/_ttsAssetForSpec:
// spec is a [applicationservices.VoiceSpec].
func (t TTSAXResourceMigrationUtilities) _ttsAssetForSpec(spec objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_ttsAssetForSpec:"), spec)
	return objectivec.Object{ID: rv}
}

// TtsAssetForSpec is an exported wrapper for the private method _ttsAssetForSpec.
func (t TTSAXResourceMigrationUtilities) TtsAssetForSpec(spec objectivec.IObject) objectivec.IObject {
	return t._ttsAssetForSpec(spec)
}
//
// spec is a [applicationservices.VoiceSpec].
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceMigrationUtilities/attributesForLegacyVoiceSpec:
// spec is a [applicationservices.VoiceSpec].
func (t TTSAXResourceMigrationUtilities) AttributesForLegacyVoiceSpec(spec objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributesForLegacyVoiceSpec:"), spec)
	return objectivec.Object{ID: rv}
}
//
// spec is a [applicationservices.VoiceSpec].
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceMigrationUtilities/getIdentifierForLegacyVoiceSpec:
// spec is a [applicationservices.VoiceSpec].
func (t TTSAXResourceMigrationUtilities) GetIdentifierForLegacyVoiceSpec(spec objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("getIdentifierForLegacyVoiceSpec:"), spec)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceMigrationUtilities/ttsIdentifierForLegacyMacIdentifier:
func (t TTSAXResourceMigrationUtilities) TtsIdentifierForLegacyMacIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ttsIdentifierForLegacyMacIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceMigrationUtilities/updatedIdentifierForLegacyIdentifier:withLanguageCode:
func (t TTSAXResourceMigrationUtilities) UpdatedIdentifierForLegacyIdentifierWithLanguageCode(identifier objectivec.IObject, code objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("updatedIdentifierForLegacyIdentifier:withLanguageCode:"), identifier, code)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceMigrationUtilities/sharedInstance
func (_TTSAXResourceMigrationUtilitiesClass TTSAXResourceMigrationUtilitiesClass) SharedInstance() TTSAXResourceMigrationUtilities {
	rv := objc.Send[objc.ID](objc.ID(_TTSAXResourceMigrationUtilitiesClass.class), objc.Sel("sharedInstance"))
	return TTSAXResourceMigrationUtilitiesFromID(rv)
}

