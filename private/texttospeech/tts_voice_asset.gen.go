// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSVoiceAsset] class.
var (
	_TTSVoiceAssetClass     TTSVoiceAssetClass
	_TTSVoiceAssetClassOnce sync.Once
)

func getTTSVoiceAssetClass() TTSVoiceAssetClass {
	_TTSVoiceAssetClassOnce.Do(func() {
		_TTSVoiceAssetClass = TTSVoiceAssetClass{class: objc.GetClass("TTSVoiceAsset")}
	})
	return _TTSVoiceAssetClass
}

// GetTTSVoiceAssetClass returns the class object for TTSVoiceAsset.
func GetTTSVoiceAssetClass() TTSVoiceAssetClass {
	return getTTSVoiceAssetClass()
}

type TTSVoiceAssetClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSVoiceAssetClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSVoiceAssetClass) Alloc() TTSVoiceAsset {
	rv := objc.Send[TTSVoiceAsset](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSVoiceAsset.DictionaryRepresentation]
//   - [TTSVoiceAsset.FileSize]
//   - [TTSVoiceAsset.SetFileSize]
//   - [TTSVoiceAsset.Footprint]
//   - [TTSVoiceAsset.Gender]
//   - [TTSVoiceAsset.Identifier]
//   - [TTSVoiceAsset.SetIdentifier]
//   - [TTSVoiceAsset.IsBuiltInVoice]
//   - [TTSVoiceAsset.IsDownloading]
//   - [TTSVoiceAsset.SetIsDownloading]
//   - [TTSVoiceAsset.IsInstalled]
//   - [TTSVoiceAsset.Languages]
//   - [TTSVoiceAsset.Name]
//   - [TTSVoiceAsset.Neural]
//   - [TTSVoiceAsset.VoicePath]
//   - [TTSVoiceAsset.SetVoicePath]
//   - [TTSVoiceAsset.VoiceType]
//   - [TTSVoiceAsset.SetVoiceType]
//   - [TTSVoiceAsset.InitWithDictionaryRepresentation]
//   - [TTSVoiceAsset.InitWithNameLanguagesGenderFootprintIsInstalledIsBuiltInMasteredVersionCompatibilityVersionNeural]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset
type TTSVoiceAsset struct {
	TTSAssetBase
}

// TTSVoiceAssetFromID constructs a [TTSVoiceAsset] from an objc.ID.
func TTSVoiceAssetFromID(id objc.ID) TTSVoiceAsset {
	return TTSVoiceAsset{TTSAssetBase: TTSAssetBaseFromID(id)}
}
// Ensure TTSVoiceAsset implements ITTSVoiceAsset.
var _ ITTSVoiceAsset = TTSVoiceAsset{}

// An interface definition for the [TTSVoiceAsset] class.
//
// # Methods
//
//   - [ITTSVoiceAsset.DictionaryRepresentation]
//   - [ITTSVoiceAsset.FileSize]
//   - [ITTSVoiceAsset.SetFileSize]
//   - [ITTSVoiceAsset.Footprint]
//   - [ITTSVoiceAsset.Gender]
//   - [ITTSVoiceAsset.Identifier]
//   - [ITTSVoiceAsset.SetIdentifier]
//   - [ITTSVoiceAsset.IsBuiltInVoice]
//   - [ITTSVoiceAsset.IsDownloading]
//   - [ITTSVoiceAsset.SetIsDownloading]
//   - [ITTSVoiceAsset.IsInstalled]
//   - [ITTSVoiceAsset.Languages]
//   - [ITTSVoiceAsset.Name]
//   - [ITTSVoiceAsset.Neural]
//   - [ITTSVoiceAsset.VoicePath]
//   - [ITTSVoiceAsset.SetVoicePath]
//   - [ITTSVoiceAsset.VoiceType]
//   - [ITTSVoiceAsset.SetVoiceType]
//   - [ITTSVoiceAsset.InitWithDictionaryRepresentation]
//   - [ITTSVoiceAsset.InitWithNameLanguagesGenderFootprintIsInstalledIsBuiltInMasteredVersionCompatibilityVersionNeural]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset
type ITTSVoiceAsset interface {
	ITTSAssetBase

	// Topic: Methods

	DictionaryRepresentation() objectivec.IObject
	FileSize() int64
	SetFileSize(value int64)
	Footprint() int64
	Gender() int64
	Identifier() string
	SetIdentifier(value string)
	IsBuiltInVoice() bool
	IsDownloading() bool
	SetIsDownloading(value bool)
	IsInstalled() bool
	Languages() foundation.INSArray
	Name() string
	Neural() bool
	VoicePath() string
	SetVoicePath(value string)
	VoiceType() int64
	SetVoiceType(value int64)
	InitWithDictionaryRepresentation(representation objectivec.IObject) TTSVoiceAsset
	InitWithNameLanguagesGenderFootprintIsInstalledIsBuiltInMasteredVersionCompatibilityVersionNeural(name objectivec.IObject, languages objectivec.IObject, gender int64, footprint int64, installed bool, in bool, version objectivec.IObject, version2 objectivec.IObject, neural bool) TTSVoiceAsset
}

// Init initializes the instance.
func (t TTSVoiceAsset) Init() TTSVoiceAsset {
	rv := objc.Send[TTSVoiceAsset](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSVoiceAsset) Autorelease() TTSVoiceAsset {
	rv := objc.Send[TTSVoiceAsset](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSVoiceAsset creates a new TTSVoiceAsset instance.
func NewTTSVoiceAsset() TTSVoiceAsset {
	class := getTTSVoiceAssetClass()
	rv := objc.Send[TTSVoiceAsset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/initWithCoder:
func NewTTSVoiceAssetWithCoder(coder objectivec.IObject) TTSVoiceAsset {
	instance := getTTSVoiceAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return TTSVoiceAssetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/initWithDictionaryRepresentation:
func NewTTSVoiceAssetWithDictionaryRepresentation(representation objectivec.IObject) TTSVoiceAsset {
	instance := getTTSVoiceAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionaryRepresentation:"), representation)
	return TTSVoiceAssetFromID(rv)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/initWithName:languages:gender:footprint:isInstalled:isBuiltIn:masteredVersion:compatibilityVersion:neural:
func NewTTSVoiceAssetWithNameLanguagesGenderFootprintIsInstalledIsBuiltInMasteredVersionCompatibilityVersionNeural(name objectivec.IObject, languages objectivec.IObject, gender int64, footprint int64, installed bool, in bool, version objectivec.IObject, version2 objectivec.IObject, neural bool) TTSVoiceAsset {
	instance := getTTSVoiceAssetClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:languages:gender:footprint:isInstalled:isBuiltIn:masteredVersion:compatibilityVersion:neural:"), name, languages, gender, footprint, installed, in, version, version2, neural)
	return TTSVoiceAssetFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/dictionaryRepresentation
func (t TTSVoiceAsset) DictionaryRepresentation() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dictionaryRepresentation"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/initWithDictionaryRepresentation:
func (t TTSVoiceAsset) InitWithDictionaryRepresentation(representation objectivec.IObject) TTSVoiceAsset {
	rv := objc.Send[TTSVoiceAsset](t.ID, objc.Sel("initWithDictionaryRepresentation:"), representation)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/initWithName:languages:gender:footprint:isInstalled:isBuiltIn:masteredVersion:compatibilityVersion:neural:
func (t TTSVoiceAsset) InitWithNameLanguagesGenderFootprintIsInstalledIsBuiltInMasteredVersionCompatibilityVersionNeural(name objectivec.IObject, languages objectivec.IObject, gender int64, footprint int64, installed bool, in bool, version objectivec.IObject, version2 objectivec.IObject, neural bool) TTSVoiceAsset {
	rv := objc.Send[TTSVoiceAsset](t.ID, objc.Sel("initWithName:languages:gender:footprint:isInstalled:isBuiltIn:masteredVersion:compatibilityVersion:neural:"), name, languages, gender, footprint, installed, in, version, version2, neural)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/fileSize
func (t TTSVoiceAsset) FileSize() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("fileSize"))
	return rv
}
func (t TTSVoiceAsset) SetFileSize(value int64) {
	objc.Send[struct{}](t.ID, objc.Sel("setFileSize:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/footprint
func (t TTSVoiceAsset) Footprint() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("footprint"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/gender
func (t TTSVoiceAsset) Gender() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("gender"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/identifier
func (t TTSVoiceAsset) Identifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSVoiceAsset) SetIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/isBuiltInVoice
func (t TTSVoiceAsset) IsBuiltInVoice() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isBuiltInVoice"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/isDownloading
func (t TTSVoiceAsset) IsDownloading() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isDownloading"))
	return rv
}
func (t TTSVoiceAsset) SetIsDownloading(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIsDownloading:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/isInstalled
func (t TTSVoiceAsset) IsInstalled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isInstalled"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/languages
func (t TTSVoiceAsset) Languages() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("languages"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/name
func (t TTSVoiceAsset) Name() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/neural
func (t TTSVoiceAsset) Neural() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("neural"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/voicePath
func (t TTSVoiceAsset) VoicePath() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voicePath"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSVoiceAsset) SetVoicePath(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoicePath:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceAsset/voiceType
func (t TTSVoiceAsset) VoiceType() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("voiceType"))
	return rv
}
func (t TTSVoiceAsset) SetVoiceType(value int64) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoiceType:"), value)
}

