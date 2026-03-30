// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAXResourceManager] class.
var (
	_TTSAXResourceManagerClass     TTSAXResourceManagerClass
	_TTSAXResourceManagerClassOnce sync.Once
)

func getTTSAXResourceManagerClass() TTSAXResourceManagerClass {
	_TTSAXResourceManagerClassOnce.Do(func() {
		_TTSAXResourceManagerClass = TTSAXResourceManagerClass{class: objc.GetClass("TTSAXResourceManager")}
	})
	return _TTSAXResourceManagerClass
}

// GetTTSAXResourceManagerClass returns the class object for TTSAXResourceManager.
func GetTTSAXResourceManagerClass() TTSAXResourceManagerClass {
	return getTTSAXResourceManagerClass()
}

type TTSAXResourceManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSAXResourceManagerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAXResourceManagerClass) Alloc() TTSAXResourceManager {
	rv := objc.Send[TTSAXResourceManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSAXResourceManager._observers]
//   - [TTSAXResourceManager.Set_observers]
//   - [TTSAXResourceManager._performBlockOnObservers]
//   - [TTSAXResourceManager._resourceWithVoiceIdAssetId]
//   - [TTSAXResourceManager._resourcesWithTypeSubTypeLanguageCode]
//   - [TTSAXResourceManager.AddObserver]
//   - [TTSAXResourceManager.AllAvailableLanguages]
//   - [TTSAXResourceManager.SetAllAvailableLanguages]
//   - [TTSAXResourceManager.AllLanguagesForVoices]
//   - [TTSAXResourceManager.AllVoices]
//   - [TTSAXResourceManager.DeleteResourceWithVoiceId]
//   - [TTSAXResourceManager.DownloadResourceWithVoiceId]
//   - [TTSAXResourceManager.DownloadResourceWithVoiceIdUserInitiated]
//   - [TTSAXResourceManager.LanguageCodeForResourceNameWithType]
//   - [TTSAXResourceManager.RemoveObserver]
//   - [TTSAXResourceManager.ResourceWithVoiceId]
//   - [TTSAXResourceManager.ResourcesWithLanguageType]
//   - [TTSAXResourceManager.ResourcesWithTypeSubType]
//   - [TTSAXResourceManager.SpeechVoiceWithVoiceId]
//   - [TTSAXResourceManager.StopDownloadResourceWithVoiceId]
//   - [TTSAXResourceManager.SuperCompactVoiceIdForCompactVoiceId]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager
type TTSAXResourceManager struct {
	objectivec.Object
}

// TTSAXResourceManagerFromID constructs a [TTSAXResourceManager] from an objc.ID.
func TTSAXResourceManagerFromID(id objc.ID) TTSAXResourceManager {
	return TTSAXResourceManager{objectivec.Object{ID: id}}
}

// Ensure TTSAXResourceManager implements ITTSAXResourceManager.
var _ ITTSAXResourceManager = TTSAXResourceManager{}

// An interface definition for the [TTSAXResourceManager] class.
//
// # Methods
//
//   - [ITTSAXResourceManager._observers]
//   - [ITTSAXResourceManager.Set_observers]
//   - [ITTSAXResourceManager._performBlockOnObservers]
//   - [ITTSAXResourceManager._resourceWithVoiceIdAssetId]
//   - [ITTSAXResourceManager._resourcesWithTypeSubTypeLanguageCode]
//   - [ITTSAXResourceManager.AddObserver]
//   - [ITTSAXResourceManager.AllAvailableLanguages]
//   - [ITTSAXResourceManager.SetAllAvailableLanguages]
//   - [ITTSAXResourceManager.AllLanguagesForVoices]
//   - [ITTSAXResourceManager.AllVoices]
//   - [ITTSAXResourceManager.DeleteResourceWithVoiceId]
//   - [ITTSAXResourceManager.DownloadResourceWithVoiceId]
//   - [ITTSAXResourceManager.DownloadResourceWithVoiceIdUserInitiated]
//   - [ITTSAXResourceManager.LanguageCodeForResourceNameWithType]
//   - [ITTSAXResourceManager.RemoveObserver]
//   - [ITTSAXResourceManager.ResourceWithVoiceId]
//   - [ITTSAXResourceManager.ResourcesWithLanguageType]
//   - [ITTSAXResourceManager.ResourcesWithTypeSubType]
//   - [ITTSAXResourceManager.SpeechVoiceWithVoiceId]
//   - [ITTSAXResourceManager.StopDownloadResourceWithVoiceId]
//   - [ITTSAXResourceManager.SuperCompactVoiceIdForCompactVoiceId]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager
type ITTSAXResourceManager interface {
	objectivec.IObject

	// Topic: Methods

	_observers() foundation.NSHashTable
	Set_observers(value foundation.NSHashTable)
	_performBlockOnObservers(observers VoidHandler)
	_resourceWithVoiceIdAssetId(id objectivec.IObject, id2 objectivec.IObject) objectivec.IObject
	_resourcesWithTypeSubTypeLanguageCode(type_ uint64, type_2 uint64, code objectivec.IObject) objectivec.IObject
	AddObserver(observer objectivec.IObject)
	AllAvailableLanguages() foundation.INSSet
	SetAllAvailableLanguages(value foundation.INSSet)
	AllLanguagesForVoices(voices bool) objectivec.IObject
	AllVoices(voices bool) objectivec.IObject
	DeleteResourceWithVoiceId(id objectivec.IObject)
	DownloadResourceWithVoiceId(id objectivec.IObject)
	DownloadResourceWithVoiceIdUserInitiated(id objectivec.IObject, initiated bool)
	LanguageCodeForResourceNameWithType(name objectivec.IObject, type_ uint64) objectivec.IObject
	RemoveObserver(observer objectivec.IObject)
	ResourceWithVoiceId(id objectivec.IObject) objectivec.IObject
	ResourcesWithLanguageType(language objectivec.IObject, type_ uint64) objectivec.IObject
	ResourcesWithTypeSubType(type_ uint64, type_2 uint64) objectivec.IObject
	SpeechVoiceWithVoiceId(id objectivec.IObject) objectivec.IObject
	StopDownloadResourceWithVoiceId(id objectivec.IObject)
	SuperCompactVoiceIdForCompactVoiceId(id objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (t TTSAXResourceManager) Init() TTSAXResourceManager {
	rv := objc.Send[TTSAXResourceManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAXResourceManager) Autorelease() TTSAXResourceManager {
	rv := objc.Send[TTSAXResourceManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAXResourceManager creates a new TTSAXResourceManager instance.
func NewTTSAXResourceManager() TTSAXResourceManager {
	class := getTTSAXResourceManagerClass()
	rv := objc.Send[TTSAXResourceManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/_performBlockOnObservers:
func (t TTSAXResourceManager) _performBlockOnObservers(observers VoidHandler) {
	_block0, _ := NewVoidBlock(observers)
	objc.Send[objc.ID](t.ID, objc.Sel("_performBlockOnObservers:"), _block0)
}

// PerformBlockOnObservers is an exported wrapper for the private method _performBlockOnObservers.
func (t TTSAXResourceManager) PerformBlockOnObservers(observers VoidHandler) {
	t._performBlockOnObservers(observers)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/_resourceWithVoiceId:assetId:
func (t TTSAXResourceManager) _resourceWithVoiceIdAssetId(id objectivec.IObject, id2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_resourceWithVoiceId:assetId:"), id, id2)
	return objectivec.Object{ID: rv}
}

// ResourceWithVoiceIdAssetId is an exported wrapper for the private method _resourceWithVoiceIdAssetId.
func (t TTSAXResourceManager) ResourceWithVoiceIdAssetId(id objectivec.IObject, id2 objectivec.IObject) objectivec.IObject {
	return t._resourceWithVoiceIdAssetId(id, id2)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/_resourcesWithType:subType:languageCode:
func (t TTSAXResourceManager) _resourcesWithTypeSubTypeLanguageCode(type_ uint64, type_2 uint64, code objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_resourcesWithType:subType:languageCode:"), type_, type_2, code)
	return objectivec.Object{ID: rv}
}

// ResourcesWithTypeSubTypeLanguageCode is an exported wrapper for the private method _resourcesWithTypeSubTypeLanguageCode.
func (t TTSAXResourceManager) ResourcesWithTypeSubTypeLanguageCode(type_ uint64, type_2 uint64, code objectivec.IObject) objectivec.IObject {
	return t._resourcesWithTypeSubTypeLanguageCode(type_, type_2, code)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/addObserver:
func (t TTSAXResourceManager) AddObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("addObserver:"), observer)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/allLanguagesForVoices:
func (t TTSAXResourceManager) AllLanguagesForVoices(voices bool) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("allLanguagesForVoices:"), voices)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/allVoices:
func (t TTSAXResourceManager) AllVoices(voices bool) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("allVoices:"), voices)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/deleteResourceWithVoiceId:
func (t TTSAXResourceManager) DeleteResourceWithVoiceId(id objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("deleteResourceWithVoiceId:"), id)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/downloadResourceWithVoiceId:
func (t TTSAXResourceManager) DownloadResourceWithVoiceId(id objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("downloadResourceWithVoiceId:"), id)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/downloadResourceWithVoiceId:userInitiated:
func (t TTSAXResourceManager) DownloadResourceWithVoiceIdUserInitiated(id objectivec.IObject, initiated bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("downloadResourceWithVoiceId:userInitiated:"), id, initiated)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/languageCodeForResourceName:withType:
func (t TTSAXResourceManager) LanguageCodeForResourceNameWithType(name objectivec.IObject, type_ uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("languageCodeForResourceName:withType:"), name, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/removeObserver:
func (t TTSAXResourceManager) RemoveObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeObserver:"), observer)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/resourceWithVoiceId:
func (t TTSAXResourceManager) ResourceWithVoiceId(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resourceWithVoiceId:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/resourcesWithLanguage:type:
func (t TTSAXResourceManager) ResourcesWithLanguageType(language objectivec.IObject, type_ uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resourcesWithLanguage:type:"), language, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/resourcesWithType:subType:
func (t TTSAXResourceManager) ResourcesWithTypeSubType(type_ uint64, type_2 uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resourcesWithType:subType:"), type_, type_2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/speechVoiceWithVoiceId:
func (t TTSAXResourceManager) SpeechVoiceWithVoiceId(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("speechVoiceWithVoiceId:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/stopDownloadResourceWithVoiceId:
func (t TTSAXResourceManager) StopDownloadResourceWithVoiceId(id objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("stopDownloadResourceWithVoiceId:"), id)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/superCompactVoiceIdForCompactVoiceId:
func (t TTSAXResourceManager) SuperCompactVoiceIdForCompactVoiceId(id objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("superCompactVoiceIdForCompactVoiceId:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/sharedInstance
func (_TTSAXResourceManagerClass TTSAXResourceManagerClass) SharedInstance() TTSAXResourceManager {
	rv := objc.Send[objc.ID](objc.ID(_TTSAXResourceManagerClass.class), objc.Sel("sharedInstance"))
	return TTSAXResourceManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/_observers
func (t TTSAXResourceManager) _observers() foundation.NSHashTable {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_observers"))
	return foundation.NSHashTableFromID(objc.ID(rv))
}
func (t TTSAXResourceManager) Set_observers(value foundation.NSHashTable) {
	objc.Send[struct{}](t.ID, objc.Sel("set_observers:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAXResourceManager/allAvailableLanguages
func (t TTSAXResourceManager) AllAvailableLanguages() foundation.INSSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("allAvailableLanguages"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (t TTSAXResourceManager) SetAllAvailableLanguages(value foundation.INSSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllAvailableLanguages:"), value)
}

// _performBlockOnObserversSync is a synchronous wrapper around [TTSAXResourceManager._performBlockOnObservers].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSAXResourceManager) _performBlockOnObserversSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t._performBlockOnObservers(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
