// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSSpeechEventObjc] class.
var (
	_TextToSpeechTTSSpeechEventObjcClass     TextToSpeechTTSSpeechEventObjcClass
	_TextToSpeechTTSSpeechEventObjcClassOnce sync.Once
)

func getTextToSpeechTTSSpeechEventObjcClass() TextToSpeechTTSSpeechEventObjcClass {
	_TextToSpeechTTSSpeechEventObjcClassOnce.Do(func() {
		_TextToSpeechTTSSpeechEventObjcClass = TextToSpeechTTSSpeechEventObjcClass{class: objc.GetClass("TextToSpeech.TTSSpeechEventObjc")}
	})
	return _TextToSpeechTTSSpeechEventObjcClass
}

// GetTextToSpeechTTSSpeechEventObjcClass returns the class object for TextToSpeech.TTSSpeechEventObjc.
func GetTextToSpeechTTSSpeechEventObjcClass() TextToSpeechTTSSpeechEventObjcClass {
	return getTextToSpeechTTSSpeechEventObjcClass()
}

type TextToSpeechTTSSpeechEventObjcClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSSpeechEventObjcClass) Alloc() TextToSpeechTTSSpeechEventObjc {
	rv := objc.Send[TextToSpeechTTSSpeechEventObjc](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSpeechEventObjc
type TextToSpeechTTSSpeechEventObjc struct {
	objectivec.Object
}

// TextToSpeechTTSSpeechEventObjcFromID constructs a [TextToSpeechTTSSpeechEventObjc] from an objc.ID.
func TextToSpeechTTSSpeechEventObjcFromID(id objc.ID) TextToSpeechTTSSpeechEventObjc {
	return TextToSpeechTTSSpeechEventObjc{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechTTSSpeechEventObjc implements ITextToSpeechTTSSpeechEventObjc.
var _ ITextToSpeechTTSSpeechEventObjc = TextToSpeechTTSSpeechEventObjc{}

// An interface definition for the [TextToSpeechTTSSpeechEventObjc] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSpeechEventObjc
type ITextToSpeechTTSSpeechEventObjc interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSSpeechEventObjc) Init() TextToSpeechTTSSpeechEventObjc {
	rv := objc.Send[TextToSpeechTTSSpeechEventObjc](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSSpeechEventObjc) Autorelease() TextToSpeechTTSSpeechEventObjc {
	rv := objc.Send[TextToSpeechTTSSpeechEventObjc](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSSpeechEventObjc creates a new TextToSpeechTTSSpeechEventObjc instance.
func NewTextToSpeechTTSSpeechEventObjc() TextToSpeechTTSSpeechEventObjc {
	class := getTextToSpeechTTSSpeechEventObjcClass()
	rv := objc.Send[TextToSpeechTTSSpeechEventObjc](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSpeechEventObjc/makeWithOtherRewrite:from:to:
func (_TextToSpeechTTSSpeechEventObjcClass TextToSpeechTTSSpeechEventObjcClass) MakeWithOtherRewriteFromTo(rewrite objectivec.IObject, from objectivec.IObject, to objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TextToSpeechTTSSpeechEventObjcClass.class), objc.Sel("makeWithOtherRewrite:from:to:"), rewrite, from, to)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSpeechEventObjc/makeWithStart:
func (_TextToSpeechTTSSpeechEventObjcClass TextToSpeechTTSSpeechEventObjcClass) MakeWithStart(start objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TextToSpeechTTSSpeechEventObjcClass.class), objc.Sel("makeWithStart:"), start)
	return objectivec.Object{ID: rv}
}

