// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSVoiceResourceAsset] class.
var (
	_TTSVoiceResourceAssetClass     TTSVoiceResourceAssetClass
	_TTSVoiceResourceAssetClassOnce sync.Once
)

func getTTSVoiceResourceAssetClass() TTSVoiceResourceAssetClass {
	_TTSVoiceResourceAssetClassOnce.Do(func() {
		_TTSVoiceResourceAssetClass = TTSVoiceResourceAssetClass{class: objc.GetClass("TTSVoiceResourceAsset")}
	})
	return _TTSVoiceResourceAssetClass
}

// GetTTSVoiceResourceAssetClass returns the class object for TTSVoiceResourceAsset.
func GetTTSVoiceResourceAssetClass() TTSVoiceResourceAssetClass {
	return getTTSVoiceResourceAssetClass()
}

type TTSVoiceResourceAssetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSVoiceResourceAssetClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSVoiceResourceAssetClass) Alloc() TTSVoiceResourceAsset {
	rv := objc.Send[TTSVoiceResourceAsset](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSVoiceResourceAsset.DefaultFootprintString]
//   - [TTSVoiceResourceAsset.DefaultTypeString]
//   - [TTSVoiceResourceAsset.DefaultVoice]
//   - [TTSVoiceResourceAsset.Languages]
//   - [TTSVoiceResourceAsset.SetLanguages]
//   - [TTSVoiceResourceAsset.ResourceList]
//   - [TTSVoiceResourceAsset.SetResourceList]
//   - [TTSVoiceResourceAsset.SearchPathURL]
//   - [TTSVoiceResourceAsset.SetSearchPathURL]
//   - [TTSVoiceResourceAsset.SyncWithConfigDataVoiceType]
//   - [TTSVoiceResourceAsset.SyncWithConfigFileVoiceType]
//   - [TTSVoiceResourceAsset.VoiceConfig]
//   - [TTSVoiceResourceAsset.SetVoiceConfig]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset
type TTSVoiceResourceAsset struct {
	TTSAssetBase
}

// TTSVoiceResourceAssetFromID constructs a [TTSVoiceResourceAsset] from an objc.ID.
func TTSVoiceResourceAssetFromID(id objc.ID) TTSVoiceResourceAsset {
	return TTSVoiceResourceAsset{TTSAssetBase: TTSAssetBaseFromID(id)}
}
// Ensure TTSVoiceResourceAsset implements ITTSVoiceResourceAsset.
var _ ITTSVoiceResourceAsset = TTSVoiceResourceAsset{}

// An interface definition for the [TTSVoiceResourceAsset] class.
//
// # Methods
//
//   - [ITTSVoiceResourceAsset.DefaultFootprintString]
//   - [ITTSVoiceResourceAsset.DefaultTypeString]
//   - [ITTSVoiceResourceAsset.DefaultVoice]
//   - [ITTSVoiceResourceAsset.Languages]
//   - [ITTSVoiceResourceAsset.SetLanguages]
//   - [ITTSVoiceResourceAsset.ResourceList]
//   - [ITTSVoiceResourceAsset.SetResourceList]
//   - [ITTSVoiceResourceAsset.SearchPathURL]
//   - [ITTSVoiceResourceAsset.SetSearchPathURL]
//   - [ITTSVoiceResourceAsset.SyncWithConfigDataVoiceType]
//   - [ITTSVoiceResourceAsset.SyncWithConfigFileVoiceType]
//   - [ITTSVoiceResourceAsset.VoiceConfig]
//   - [ITTSVoiceResourceAsset.SetVoiceConfig]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset
type ITTSVoiceResourceAsset interface {
	ITTSAssetBase

	// Topic: Methods

	DefaultFootprintString() objectivec.IObject
	DefaultTypeString() objectivec.IObject
	DefaultVoice() objectivec.IObject
	Languages() foundation.INSArray
	SetLanguages(value foundation.INSArray)
	ResourceList() foundation.INSArray
	SetResourceList(value foundation.INSArray)
	SearchPathURL() foundation.INSURL
	SetSearchPathURL(value foundation.INSURL)
	SyncWithConfigDataVoiceType(data objectivec.IObject, type_ int64)
	SyncWithConfigFileVoiceType(file objectivec.IObject, type_ int64)
	VoiceConfig() foundation.INSDictionary
	SetVoiceConfig(value foundation.INSDictionary)
}

// Init initializes the instance.
func (t TTSVoiceResourceAsset) Init() TTSVoiceResourceAsset {
	rv := objc.Send[TTSVoiceResourceAsset](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSVoiceResourceAsset) Autorelease() TTSVoiceResourceAsset {
	rv := objc.Send[TTSVoiceResourceAsset](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSVoiceResourceAsset creates a new TTSVoiceResourceAsset instance.
func NewTTSVoiceResourceAsset() TTSVoiceResourceAsset {
	class := getTTSVoiceResourceAssetClass()
	rv := objc.Send[TTSVoiceResourceAsset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/initWithCoder:
func NewTTSVoiceResourceAssetWithCoder(coder objectivec.IObject) TTSVoiceResourceAsset {
	instance := getTTSVoiceResourceAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return TTSVoiceResourceAssetFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/defaultFootprintString
func (t TTSVoiceResourceAsset) DefaultFootprintString() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("defaultFootprintString"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/defaultTypeString
func (t TTSVoiceResourceAsset) DefaultTypeString() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("defaultTypeString"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/defaultVoice
func (t TTSVoiceResourceAsset) DefaultVoice() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("defaultVoice"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/syncWithConfigData:voiceType:
func (t TTSVoiceResourceAsset) SyncWithConfigDataVoiceType(data objectivec.IObject, type_ int64) {
	objc.Send[objc.ID](t.ID, objc.Sel("syncWithConfigData:voiceType:"), data, type_)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/syncWithConfigFile:voiceType:
func (t TTSVoiceResourceAsset) SyncWithConfigFileVoiceType(file objectivec.IObject, type_ int64) {
	objc.Send[objc.ID](t.ID, objc.Sel("syncWithConfigFile:voiceType:"), file, type_)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/legacyPlatforms
func (_TTSVoiceResourceAssetClass TTSVoiceResourceAssetClass) LegacyPlatforms() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSVoiceResourceAssetClass.class), objc.Sel("legacyPlatforms"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/languages
func (t TTSVoiceResourceAsset) Languages() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("languages"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSVoiceResourceAsset) SetLanguages(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setLanguages:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/resourceList
func (t TTSVoiceResourceAsset) ResourceList() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resourceList"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSVoiceResourceAsset) SetResourceList(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setResourceList:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/searchPathURL
func (t TTSVoiceResourceAsset) SearchPathURL() foundation.INSURL {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("searchPathURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (t TTSVoiceResourceAsset) SetSearchPathURL(value foundation.INSURL) {
	objc.Send[struct{}](t.ID, objc.Sel("setSearchPathURL:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceAsset/voiceConfig
func (t TTSVoiceResourceAsset) VoiceConfig() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceConfig"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSVoiceResourceAsset) SetVoiceConfig(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoiceConfig:"), value)
}

