// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSSSEUtils] class.
var (
	_TTSSSEUtilsClass     TTSSSEUtilsClass
	_TTSSSEUtilsClassOnce sync.Once
)

func getTTSSSEUtilsClass() TTSSSEUtilsClass {
	_TTSSSEUtilsClassOnce.Do(func() {
		_TTSSSEUtilsClass = TTSSSEUtilsClass{class: objc.GetClass("TTSSSEUtils")}
	})
	return _TTSSSEUtilsClass
}

// GetTTSSSEUtilsClass returns the class object for TTSSSEUtils.
func GetTTSSSEUtilsClass() TTSSSEUtilsClass {
	return getTTSSSEUtilsClass()
}

type TTSSSEUtilsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSSEUtilsClass) Alloc() TTSSSEUtils {
	rv := objc.Send[TTSSSEUtils](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSSEUtils
type TTSSSEUtils struct {
	objectivec.Object
}

// TTSSSEUtilsFromID constructs a [TTSSSEUtils] from an objc.ID.
func TTSSSEUtilsFromID(id objc.ID) TTSSSEUtils {
	return TTSSSEUtils{objectivec.Object{ID: id}}
}
// Ensure TTSSSEUtils implements ITTSSSEUtils.
var _ ITTSSSEUtils = TTSSSEUtils{}

// An interface definition for the [TTSSSEUtils] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSSEUtils
type ITTSSSEUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSSSEUtils) Init() TTSSSEUtils {
	rv := objc.Send[TTSSSEUtils](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSSEUtils) Autorelease() TTSSSEUtils {
	rv := objc.Send[TTSSSEUtils](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSSEUtils creates a new TTSSSEUtils instance.
func NewTTSSSEUtils() TTSSSEUtils {
	class := getTTSSSEUtilsClass()
	rv := objc.Send[TTSSSEUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSSEUtils/combinedProsodyMarkupForString:rate:pitch:volume:
func (_TTSSSEUtilsClass TTSSSEUtilsClass) CombinedProsodyMarkupForStringRatePitchVolume(string_ objectivec.IObject, rate objectivec.IObject, pitch objectivec.IObject, volume objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSSEUtilsClass.class), objc.Sel("combinedProsodyMarkupForString:rate:pitch:volume:"), string_, rate, pitch, volume)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSSEUtils/enclosedStringWithPhonemes:originalString:
func (_TTSSSEUtilsClass TTSSSEUtilsClass) EnclosedStringWithPhonemesOriginalString(phonemes objectivec.IObject, string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSSEUtilsClass.class), objc.Sel("enclosedStringWithPhonemes:originalString:"), phonemes, string_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSSEUtils/genericMarkerMarkupWithName:
func (_TTSSSEUtilsClass TTSSSEUtilsClass) GenericMarkerMarkupWithName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSSEUtilsClass.class), objc.Sel("genericMarkerMarkupWithName:"), name)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSSEUtils/speechMarkupStringForType:string:
func (_TTSSSEUtilsClass TTSSSEUtilsClass) SpeechMarkupStringForTypeString(type_ int64, string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSSEUtilsClass.class), objc.Sel("speechMarkupStringForType:string:"), type_, string_)
	return objectivec.Object{ID: rv}
}

