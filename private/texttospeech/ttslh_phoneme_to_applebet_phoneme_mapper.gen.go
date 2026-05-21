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

func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _acceptRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_acceptRule"))
}

// AcceptRule is an exported wrapper for the private method _acceptRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) AcceptRule() error {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_acceptRule")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_acceptRule"}
		return err
	}
	_TTSLHPhonemeToApplebetPhonemeMapperClass._acceptRule()
	return nil
}

// CanAcceptRule reports whether the receiver responds to the private selector _acceptRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanAcceptRule() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_acceptRule"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _initializeRules() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_initializeRules"))
	return objectivec.Object{ID: rv}
}

// InitializeRules is an exported wrapper for the private method _initializeRules.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) InitializeRules() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_initializeRules")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initializeRules"}
		return nil, err
	}
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._initializeRules(), nil
}

// CanInitializeRules reports whether the receiver responds to the private selector _initializeRules.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanInitializeRules() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_initializeRules"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _leftRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_leftRaisingContextRule"))
}

// LeftRaisingContextRule is an exported wrapper for the private method _leftRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) LeftRaisingContextRule() error {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_leftRaisingContextRule")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_leftRaisingContextRule"}
		return err
	}
	_TTSLHPhonemeToApplebetPhonemeMapperClass._leftRaisingContextRule()
	return nil
}

// CanLeftRaisingContextRule reports whether the receiver responds to the private selector _leftRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanLeftRaisingContextRule() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_leftRaisingContextRule"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _phonemeArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonemeArray:"), array)
	return objectivec.Object{ID: rv}
}

// PhonemeArray is an exported wrapper for the private method _phonemeArray.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) PhonemeArray(array objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonemeArray:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_phonemeArray:"}
		return nil, err
	}
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._phonemeArray(array), nil
}

// CanPhonemeArray reports whether the receiver responds to the private selector _phonemeArray:.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanPhonemeArray() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonemeArray:"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _phonemeRules() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonemeRules"))
	return objectivec.Object{ID: rv}
}

// PhonemeRules is an exported wrapper for the private method _phonemeRules.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) PhonemeRules() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonemeRules")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_phonemeRules"}
		return nil, err
	}
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._phonemeRules(), nil
}

// CanPhonemeRules reports whether the receiver responds to the private selector _phonemeRules.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanPhonemeRules() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonemeRules"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _phonoMatchMatchMatchposCount(match objectivec.IObject, match2 objectivec.IObject, matchpos []int, count int) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonoMatch:match:matchpos:count:"), match, match2, objc.CArray(matchpos), count)
	return objectivec.Object{ID: rv}
}

// PhonoMatchMatchMatchposCount is an exported wrapper for the private method _phonoMatchMatchMatchposCount.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) PhonoMatchMatchMatchposCount(match objectivec.IObject, match2 objectivec.IObject, matchpos []int, count int) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonoMatch:match:matchpos:count:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_phonoMatch:match:matchpos:count:"}
		return nil, err
	}
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._phonoMatchMatchMatchposCount(match, match2, matchpos, count), nil
}

// CanPhonoMatchMatchMatchposCount reports whether the receiver responds to the private selector _phonoMatch:match:matchpos:count:.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanPhonoMatchMatchMatchposCount() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonoMatch:match:matchpos:count:"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _phonoTranslation(translation objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonoTranslation:"), translation)
	return objectivec.Object{ID: rv}
}

// PhonoTranslation is an exported wrapper for the private method _phonoTranslation.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) PhonoTranslation(translation objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonoTranslation:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_phonoTranslation:"}
		return nil, err
	}
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._phonoTranslation(translation), nil
}

// CanPhonoTranslation reports whether the receiver responds to the private selector _phonoTranslation:.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanPhonoTranslation() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_phonoTranslation:"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _retrieveRegularExpression(expression objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_retrieveRegularExpression:"), expression)
	return objectivec.Object{ID: rv}
}

// RetrieveRegularExpression is an exported wrapper for the private method _retrieveRegularExpression.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) RetrieveRegularExpression(expression objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_retrieveRegularExpression:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_retrieveRegularExpression:"}
		return nil, err
	}
	return _TTSLHPhonemeToApplebetPhonemeMapperClass._retrieveRegularExpression(expression), nil
}

// CanRetrieveRegularExpression reports whether the receiver responds to the private selector _retrieveRegularExpression:.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanRetrieveRegularExpression() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_retrieveRegularExpression:"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _rightRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_rightRaisingContextRule"))
}

// RightRaisingContextRule is an exported wrapper for the private method _rightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) RightRaisingContextRule() error {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_rightRaisingContextRule")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_rightRaisingContextRule"}
		return err
	}
	_TTSLHPhonemeToApplebetPhonemeMapperClass._rightRaisingContextRule()
	return nil
}

// CanRightRaisingContextRule reports whether the receiver responds to the private selector _rightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanRightRaisingContextRule() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_rightRaisingContextRule"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _uberLeftRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberLeftRaisingContextRule"))
}

// UberLeftRaisingContextRule is an exported wrapper for the private method _uberLeftRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) UberLeftRaisingContextRule() error {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberLeftRaisingContextRule")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_uberLeftRaisingContextRule"}
		return err
	}
	_TTSLHPhonemeToApplebetPhonemeMapperClass._uberLeftRaisingContextRule()
	return nil
}

// CanUberLeftRaisingContextRule reports whether the receiver responds to the private selector _uberLeftRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanUberLeftRaisingContextRule() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberLeftRaisingContextRule"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _uberRightRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberRightRaisingContextRule"))
}

// UberRightRaisingContextRule is an exported wrapper for the private method _uberRightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) UberRightRaisingContextRule() error {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberRightRaisingContextRule")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_uberRightRaisingContextRule"}
		return err
	}
	_TTSLHPhonemeToApplebetPhonemeMapperClass._uberRightRaisingContextRule()
	return nil
}

// CanUberRightRaisingContextRule reports whether the receiver responds to the private selector _uberRightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanUberRightRaisingContextRule() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberRightRaisingContextRule"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _uberUberRightRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberUberRightRaisingContextRule"))
}

// UberUberRightRaisingContextRule is an exported wrapper for the private method _uberUberRightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) UberUberRightRaisingContextRule() error {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberUberRightRaisingContextRule")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_uberUberRightRaisingContextRule"}
		return err
	}
	_TTSLHPhonemeToApplebetPhonemeMapperClass._uberUberRightRaisingContextRule()
	return nil
}

// CanUberUberRightRaisingContextRule reports whether the receiver responds to the private selector _uberUberRightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanUberUberRightRaisingContextRule() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberUberRightRaisingContextRule"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) _uberUberUberRightRaisingContextRule() {
	objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberUberUberRightRaisingContextRule"))
}

// UberUberUberRightRaisingContextRule is an exported wrapper for the private method _uberUberUberRightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) UberUberUberRightRaisingContextRule() error {
	if !objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberUberUberRightRaisingContextRule")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_uberUberUberRightRaisingContextRule"}
		return err
	}
	_TTSLHPhonemeToApplebetPhonemeMapperClass._uberUberUberRightRaisingContextRule()
	return nil
}

// CanUberUberUberRightRaisingContextRule reports whether the receiver responds to the private selector _uberUberUberRightRaisingContextRule.
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) CanUberUberUberRightRaisingContextRule() bool {
	return objc.RespondsToSelector(objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("_uberUberUberRightRaisingContextRule"))
}
func (_TTSLHPhonemeToApplebetPhonemeMapperClass TTSLHPhonemeToApplebetPhonemeMapperClass) ConvertLHToApplebet(applebet objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSLHPhonemeToApplebetPhonemeMapperClass.class), objc.Sel("convertLHToApplebet:"), applebet)
	return objectivec.Object{ID: rv}
}
