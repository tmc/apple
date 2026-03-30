// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSEmojiUtilities] class.
var (
	_TTSEmojiUtilitiesClass     TTSEmojiUtilitiesClass
	_TTSEmojiUtilitiesClassOnce sync.Once
)

func getTTSEmojiUtilitiesClass() TTSEmojiUtilitiesClass {
	_TTSEmojiUtilitiesClassOnce.Do(func() {
		_TTSEmojiUtilitiesClass = TTSEmojiUtilitiesClass{class: objc.GetClass("TTSEmojiUtilities")}
	})
	return _TTSEmojiUtilitiesClass
}

// GetTTSEmojiUtilitiesClass returns the class object for TTSEmojiUtilities.
func GetTTSEmojiUtilitiesClass() TTSEmojiUtilitiesClass {
	return getTTSEmojiUtilitiesClass()
}

type TTSEmojiUtilitiesClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSEmojiUtilitiesClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSEmojiUtilitiesClass) Alloc() TTSEmojiUtilities {
	rv := objc.Send[TTSEmojiUtilities](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSEmojiUtilities
type TTSEmojiUtilities struct {
	objectivec.Object
}

// TTSEmojiUtilitiesFromID constructs a [TTSEmojiUtilities] from an objc.ID.
func TTSEmojiUtilitiesFromID(id objc.ID) TTSEmojiUtilities {
	return TTSEmojiUtilities{objectivec.Object{ID: id}}
}

// Ensure TTSEmojiUtilities implements ITTSEmojiUtilities.
var _ ITTSEmojiUtilities = TTSEmojiUtilities{}

// An interface definition for the [TTSEmojiUtilities] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSEmojiUtilities
type ITTSEmojiUtilities interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSEmojiUtilities) Init() TTSEmojiUtilities {
	rv := objc.Send[TTSEmojiUtilities](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSEmojiUtilities) Autorelease() TTSEmojiUtilities {
	rv := objc.Send[TTSEmojiUtilities](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSEmojiUtilities creates a new TTSEmojiUtilities instance.
func NewTTSEmojiUtilities() TTSEmojiUtilities {
	class := getTTSEmojiUtilitiesClass()
	rv := objc.Send[TTSEmojiUtilities](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSEmojiUtilities/_initializeEmojiStructures:
func (_TTSEmojiUtilitiesClass TTSEmojiUtilitiesClass) _initializeEmojiStructures(structures objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_TTSEmojiUtilitiesClass.class), objc.Sel("_initializeEmojiStructures:"), structures)
}

// InitializeEmojiStructures is an exported wrapper for the private method _initializeEmojiStructures.
func (_TTSEmojiUtilitiesClass TTSEmojiUtilitiesClass) InitializeEmojiStructures(structures objectivec.IObject) {
	_TTSEmojiUtilitiesClass._initializeEmojiStructures(structures)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSEmojiUtilities/emojiRangeFromString:withSearchRange:
func (_TTSEmojiUtilitiesClass TTSEmojiUtilitiesClass) EmojiRangeFromStringWithSearchRange(string_ objectivec.IObject, range_ foundation.NSRange) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](objc.ID(_TTSEmojiUtilitiesClass.class), objc.Sel("emojiRangeFromString:withSearchRange:"), string_, range_)
	return foundation.NSRange(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSEmojiUtilities/enumerateEmojiCharactersInString:languageCode:withBlock:
func (_TTSEmojiUtilitiesClass TTSEmojiUtilitiesClass) EnumerateEmojiCharactersInStringLanguageCodeWithBlock(string_ objectivec.IObject, code objectivec.IObject, block VoidHandler) {
	_block2, _ := NewVoidBlock(block)
	objc.Send[objc.ID](objc.ID(_TTSEmojiUtilitiesClass.class), objc.Sel("enumerateEmojiCharactersInString:languageCode:withBlock:"), string_, code, _block2)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSEmojiUtilities/stringByRemovingEmojiCharacters:
func (_TTSEmojiUtilitiesClass TTSEmojiUtilitiesClass) StringByRemovingEmojiCharacters(characters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSEmojiUtilitiesClass.class), objc.Sel("stringByRemovingEmojiCharacters:"), characters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSEmojiUtilities/stringByReplacingEmojiCharactersWithEmojiDescriptions:stringForPauses:language:rangeReplacements:appendEmojiSuffix:
func (_TTSEmojiUtilitiesClass TTSEmojiUtilitiesClass) StringByReplacingEmojiCharactersWithEmojiDescriptionsStringForPausesLanguageRangeReplacementsAppendEmojiSuffix(descriptions objectivec.IObject, pauses objectivec.IObject, language objectivec.IObject, replacements objectivec.IObject, suffix bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSEmojiUtilitiesClass.class), objc.Sel("stringByReplacingEmojiCharactersWithEmojiDescriptions:stringForPauses:language:rangeReplacements:appendEmojiSuffix:"), descriptions, pauses, language, replacements, suffix)
	return objectivec.Object{ID: rv}
}

// EnumerateEmojiCharactersInStringLanguageCodeWithBlockSync is a synchronous wrapper around [TTSEmojiUtilities.EnumerateEmojiCharactersInStringLanguageCodeWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (tc TTSEmojiUtilitiesClass) EnumerateEmojiCharactersInStringLanguageCodeWithBlockSync(ctx context.Context, string_ objectivec.IObject, code objectivec.IObject) error {
	done := make(chan struct{}, 1)
	tc.EnumerateEmojiCharactersInStringLanguageCodeWithBlock(string_, code, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
