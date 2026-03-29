// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSLHPhonemeToApplebetPhonemeMapper] class.
var (
	_TTSLHPhonemeToApplebetPhonemeMapperClass     TTSLHPhonemeToApplebetPhonemeMapperClass
	_TTSLHPhonemeToApplebetPhonemeMapperClassOnce sync.Once
)

func getTTSLHPhonemeToApplebetPhonemeMapperClass() TTSLHPhonemeToApplebetPhonemeMapperClass {
	_TTSLHPhonemeToApplebetPhonemeMapperClassOnce.Do(func() {
		_TTSLHPhonemeToApplebetPhonemeMapperClass = TTSLHPhonemeToApplebetPhonemeMapperClass{class: objc.GetClass("TTSLHPhonemeToApplebetPhonemeMapper")}
	})
	return _TTSLHPhonemeToApplebetPhonemeMapperClass
}

// GetTTSLHPhonemeToApplebetPhonemeMapperClass returns the class object for TTSLHPhonemeToApplebetPhonemeMapper.
func GetTTSLHPhonemeToApplebetPhonemeMapperClass() TTSLHPhonemeToApplebetPhonemeMapperClass {
	return getTTSLHPhonemeToApplebetPhonemeMapperClass()
}

type TTSLHPhonemeToApplebetPhonemeMapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSLHPhonemeToApplebetPhonemeMapperClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSLHPhonemeToApplebetPhonemeMapperClass) Alloc() TTSLHPhonemeToApplebetPhonemeMapper {
	rv := objc.Send[TTSLHPhonemeToApplebetPhonemeMapper](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper
type TTSLHPhonemeToApplebetPhonemeMapper struct {
	objectivec.Object
}

// TTSLHPhonemeToApplebetPhonemeMapperFromID constructs a [TTSLHPhonemeToApplebetPhonemeMapper] from an objc.ID.
func TTSLHPhonemeToApplebetPhonemeMapperFromID(id objc.ID) TTSLHPhonemeToApplebetPhonemeMapper {
	return TTSLHPhonemeToApplebetPhonemeMapper{objectivec.Object{ID: id}}
}
// Ensure TTSLHPhonemeToApplebetPhonemeMapper implements ITTSLHPhonemeToApplebetPhonemeMapper.
var _ ITTSLHPhonemeToApplebetPhonemeMapper = TTSLHPhonemeToApplebetPhonemeMapper{}

// An interface definition for the [TTSLHPhonemeToApplebetPhonemeMapper] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper
type ITTSLHPhonemeToApplebetPhonemeMapper interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSLHPhonemeToApplebetPhonemeMapper) Init() TTSLHPhonemeToApplebetPhonemeMapper {
	rv := objc.Send[TTSLHPhonemeToApplebetPhonemeMapper](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSLHPhonemeToApplebetPhonemeMapper) Autorelease() TTSLHPhonemeToApplebetPhonemeMapper {
	rv := objc.Send[TTSLHPhonemeToApplebetPhonemeMapper](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSLHPhonemeToApplebetPhonemeMapper creates a new TTSLHPhonemeToApplebetPhonemeMapper instance.
func NewTTSLHPhonemeToApplebetPhonemeMapper() TTSLHPhonemeToApplebetPhonemeMapper {
	class := getTTSLHPhonemeToApplebetPhonemeMapperClass()
	rv := objc.Send[TTSLHPhonemeToApplebetPhonemeMapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_acceptRule
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _acceptRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_acceptRule"))
}

// AcceptRule is an exported wrapper for the private method _acceptRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) AcceptRule() {
	_TTSLHPhonemeToApplebetPhonemeMapperClass._acceptRule()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_initializeRules
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _initializeRules() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_initializeRules"))
	return objectivec.Object{ID: rv}
}

// InitializeRules is an exported wrapper for the private method _initializeRules.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) InitializeRules() objectivec.IObject {
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._initializeRules()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_leftRaisingContextRule
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _leftRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_leftRaisingContextRule"))
}

// LeftRaisingContextRule is an exported wrapper for the private method _leftRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) LeftRaisingContextRule() {
	_TTSLHPhonemeToApplebetPhonemeMapperClass._leftRaisingContextRule()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_phonemeArray:
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _phonemeArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonemeArray:"), array)
	return objectivec.Object{ID: rv}
}

// PhonemeArray is an exported wrapper for the private method _phonemeArray.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) PhonemeArray(array objectivec.IObject) objectivec.IObject {
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._phonemeArray(array)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_phonemeRules
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _phonemeRules() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonemeRules"))
	return objectivec.Object{ID: rv}
}

// PhonemeRules is an exported wrapper for the private method _phonemeRules.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) PhonemeRules() objectivec.IObject {
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._phonemeRules()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_phonoMatch:match:matchpos:count:
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _phonoMatchMatchMatchposCount(match objectivec.IObject, match2 objectivec.IObject, matchpos []int, count int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonoMatch:match:matchpos:count:"), match, match2, objc.CArray(matchpos), count)
	return objectivec.Object{ID: rv}
}

// PhonoMatchMatchMatchposCount is an exported wrapper for the private method _phonoMatchMatchMatchposCount.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) PhonoMatchMatchMatchposCount(match objectivec.IObject, match2 objectivec.IObject, matchpos []int, count int) objectivec.IObject {
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._phonoMatchMatchMatchposCount(match, match2, matchpos, count)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_phonoTranslation:
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _phonoTranslation(translation objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonoTranslation:"), translation)
	return objectivec.Object{ID: rv}
}

// PhonoTranslation is an exported wrapper for the private method _phonoTranslation.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) PhonoTranslation(translation objectivec.IObject) objectivec.IObject {
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._phonoTranslation(translation)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_retrieveRegularExpression:
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _retrieveRegularExpression(expression objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_retrieveRegularExpression:"), expression)
	return objectivec.Object{ID: rv}
}

// RetrieveRegularExpression is an exported wrapper for the private method _retrieveRegularExpression.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) RetrieveRegularExpression(expression objectivec.IObject) objectivec.IObject {
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._retrieveRegularExpression(expression)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_rightRaisingContextRule
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _rightRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_rightRaisingContextRule"))
}

// RightRaisingContextRule is an exported wrapper for the private method _rightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) RightRaisingContextRule() {
	_TTSLHPhonemeToApplebetPhonemeMapperClass._rightRaisingContextRule()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_uberLeftRaisingContextRule
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _uberLeftRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberLeftRaisingContextRule"))
}

// UberLeftRaisingContextRule is an exported wrapper for the private method _uberLeftRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) UberLeftRaisingContextRule() {
	_TTSLHPhonemeToApplebetPhonemeMapperClass._uberLeftRaisingContextRule()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_uberRightRaisingContextRule
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _uberRightRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberRightRaisingContextRule"))
}

// UberRightRaisingContextRule is an exported wrapper for the private method _uberRightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) UberRightRaisingContextRule() {
	_TTSLHPhonemeToApplebetPhonemeMapperClass._uberRightRaisingContextRule()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_uberUberRightRaisingContextRule
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _uberUberRightRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberUberRightRaisingContextRule"))
}

// UberUberRightRaisingContextRule is an exported wrapper for the private method _uberUberRightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) UberUberRightRaisingContextRule() {
	_TTSLHPhonemeToApplebetPhonemeMapperClass._uberUberRightRaisingContextRule()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/_uberUberUberRightRaisingContextRule
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _uberUberUberRightRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberUberUberRightRaisingContextRule"))
}

// UberUberUberRightRaisingContextRule is an exported wrapper for the private method _uberUberUberRightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) UberUberUberRightRaisingContextRule() {
	_TTSLHPhonemeToApplebetPhonemeMapperClass._uberUberUberRightRaisingContextRule()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSLHPhonemeToApplebetPhonemeMapper/convertLHToApplebet:
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) ConvertLHToApplebet(applebet objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("convertLHToApplebet:"), applebet)
	return objectivec.Object{ID: rv}
}

