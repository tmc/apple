// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[TTSVoiceResourceAsset](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSVoiceResourceAsset.Languages]
//   - [TTSVoiceResourceAsset.SetLanguages]
//   - [TTSVoiceResourceAsset.ResourceList]
//   - [TTSVoiceResourceAsset.SetResourceList]
//   - [TTSVoiceResourceAsset.SearchPathURL]
//   - [TTSVoiceResourceAsset.SetSearchPathURL]
//   - [TTSVoiceResourceAsset.SyncWithConfigData]
//   - [TTSVoiceResourceAsset.SyncWithConfigFile]
//   - [TTSVoiceResourceAsset.VoiceConfig]
//   - [TTSVoiceResourceAsset.SetVoiceConfig]
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
//   - [ITTSVoiceResourceAsset.Languages]
//   - [ITTSVoiceResourceAsset.SetLanguages]
//   - [ITTSVoiceResourceAsset.ResourceList]
//   - [ITTSVoiceResourceAsset.SetResourceList]
//   - [ITTSVoiceResourceAsset.SearchPathURL]
//   - [ITTSVoiceResourceAsset.SetSearchPathURL]
//   - [ITTSVoiceResourceAsset.SyncWithConfigData]
//   - [ITTSVoiceResourceAsset.SyncWithConfigFile]
//   - [ITTSVoiceResourceAsset.VoiceConfig]
//   - [ITTSVoiceResourceAsset.SetVoiceConfig]
type ITTSVoiceResourceAsset interface {
	ITTSAssetBase

	// Topic: Methods

	Languages() foundation.INSArray
	SetLanguages(value foundation.INSArray)
	ResourceList() foundation.INSArray
	SetResourceList(value foundation.INSArray)
	SearchPathURL() foundation.NSURL
	SetSearchPathURL(value foundation.NSURL)
	SyncWithConfigData(data objectivec.IObject)
	SyncWithConfigFile(file objectivec.IObject)
	VoiceConfig() foundation.INSDictionary
	SetVoiceConfig(value foundation.INSDictionary)
}

// Init initializes the instance.
func (t TTSVoiceResourceAsset) Init() TTSVoiceResourceAsset {
	rv := objc.SendIfResponds[TTSVoiceResourceAsset](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSVoiceResourceAsset) Autorelease() TTSVoiceResourceAsset {
	rv := objc.SendIfResponds[TTSVoiceResourceAsset](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSVoiceResourceAsset creates a new TTSVoiceResourceAsset instance.
func NewTTSVoiceResourceAsset() TTSVoiceResourceAsset {
	class := getTTSVoiceResourceAssetClass()
	rv := objc.SendIfResponds[TTSVoiceResourceAsset](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTTSVoiceResourceAssetWithCoder(coder objectivec.IObject) TTSVoiceResourceAsset {
	instance := getTTSVoiceResourceAssetClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return TTSVoiceResourceAssetFromID(rv)
}

func (t TTSVoiceResourceAsset) SyncWithConfigData(data objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("syncWithConfigData:"), data)
}
func (t TTSVoiceResourceAsset) SyncWithConfigFile(file objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("syncWithConfigFile:"), file)
}

func (t TTSVoiceResourceAsset) Languages() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("languages"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSVoiceResourceAsset) SetLanguages(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setLanguages:"), value)
}
func (t TTSVoiceResourceAsset) ResourceList() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("resourceList"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSVoiceResourceAsset) SetResourceList(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setResourceList:"), value)
}
func (t TTSVoiceResourceAsset) SearchPathURL() foundation.NSURL {
	rv := objc.SendIfResponds[foundation.NSURL](t.ID, objc.Sel("searchPathURL"))
	return foundation.NSURL(rv)
}
func (t TTSVoiceResourceAsset) SetSearchPathURL(value foundation.NSURL) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setSearchPathURL:"), value)
}
func (t TTSVoiceResourceAsset) VoiceConfig() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("voiceConfig"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSVoiceResourceAsset) SetVoiceConfig(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setVoiceConfig:"), value)
}
