// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSVoiceResourceManager] class.
var (
	_TTSVoiceResourceManagerClass     TTSVoiceResourceManagerClass
	_TTSVoiceResourceManagerClassOnce sync.Once
)

func getTTSVoiceResourceManagerClass() TTSVoiceResourceManagerClass {
	_TTSVoiceResourceManagerClassOnce.Do(func() {
		_TTSVoiceResourceManagerClass = TTSVoiceResourceManagerClass{class: objc.GetClass("TTSVoiceResourceManager")}
	})
	return _TTSVoiceResourceManagerClass
}

// GetTTSVoiceResourceManagerClass returns the class object for TTSVoiceResourceManager.
func GetTTSVoiceResourceManagerClass() TTSVoiceResourceManagerClass {
	return getTTSVoiceResourceManagerClass()
}

type TTSVoiceResourceManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSVoiceResourceManagerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSVoiceResourceManagerClass) Alloc() TTSVoiceResourceManager {
	rv := objc.Send[TTSVoiceResourceManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceManager
type TTSVoiceResourceManager struct {
	objectivec.Object
}

// TTSVoiceResourceManagerFromID constructs a [TTSVoiceResourceManager] from an objc.ID.
func TTSVoiceResourceManagerFromID(id objc.ID) TTSVoiceResourceManager {
	return TTSVoiceResourceManager{objectivec.Object{ID: id}}
}

// Ensure TTSVoiceResourceManager implements ITTSVoiceResourceManager.
var _ ITTSVoiceResourceManager = TTSVoiceResourceManager{}

// An interface definition for the [TTSVoiceResourceManager] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceManager
type ITTSVoiceResourceManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSVoiceResourceManager) Init() TTSVoiceResourceManager {
	rv := objc.Send[TTSVoiceResourceManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSVoiceResourceManager) Autorelease() TTSVoiceResourceManager {
	rv := objc.Send[TTSVoiceResourceManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSVoiceResourceManager creates a new TTSVoiceResourceManager instance.
func NewTTSVoiceResourceManager() TTSVoiceResourceManager {
	class := getTTSVoiceResourceManagerClass()
	rv := objc.Send[TTSVoiceResourceManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceManager/effectiveResourceForLanguageCode:andVoiceType:
func (_TTSVoiceResourceManagerClass TTSVoiceResourceManagerClass) EffectiveResourceForLanguageCodeAndVoiceType(code objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSVoiceResourceManagerClass.class), objc.Sel("effectiveResourceForLanguageCode:andVoiceType:"), code, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSVoiceResourceManager/enumerateLoadableResourcesInAsset:usingBlock:
func (_TTSVoiceResourceManagerClass TTSVoiceResourceManagerClass) EnumerateLoadableResourcesInAssetUsingBlock(asset objectivec.IObject, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](objc.ID(_TTSVoiceResourceManagerClass.class), objc.Sel("enumerateLoadableResourcesInAsset:usingBlock:"), asset, _block1)
}

// EnumerateLoadableResourcesInAssetUsingBlockSync is a synchronous wrapper around [TTSVoiceResourceManager.EnumerateLoadableResourcesInAssetUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (tc TTSVoiceResourceManagerClass) EnumerateLoadableResourcesInAssetUsingBlockSync(ctx context.Context, asset objectivec.IObject) error {
	done := make(chan struct{}, 1)
	tc.EnumerateLoadableResourcesInAssetUsingBlock(asset, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
