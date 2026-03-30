// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechDebouncer] class.
var (
	_TextToSpeechDebouncerClass     TextToSpeechDebouncerClass
	_TextToSpeechDebouncerClassOnce sync.Once
)

func getTextToSpeechDebouncerClass() TextToSpeechDebouncerClass {
	_TextToSpeechDebouncerClassOnce.Do(func() {
		_TextToSpeechDebouncerClass = TextToSpeechDebouncerClass{class: objc.GetClass("TextToSpeech.Debouncer")}
	})
	return _TextToSpeechDebouncerClass
}

// GetTextToSpeechDebouncerClass returns the class object for TextToSpeech.Debouncer.
func GetTextToSpeechDebouncerClass() TextToSpeechDebouncerClass {
	return getTextToSpeechDebouncerClass()
}

type TextToSpeechDebouncerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechDebouncerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechDebouncerClass) Alloc() TextToSpeechDebouncer {
	rv := objc.Send[TextToSpeechDebouncer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.Debouncer
type TextToSpeechDebouncer struct {
	objectivec.Object
}

// TextToSpeechDebouncerFromID constructs a [TextToSpeechDebouncer] from an objc.ID.
func TextToSpeechDebouncerFromID(id objc.ID) TextToSpeechDebouncer {
	return TextToSpeechDebouncer{objectivec.Object{ID: id}}
}

// NOTE: TextToSpeechDebouncer struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechDebouncer embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechDebouncer] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.Debouncer
type ITextToSpeechDebouncer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechDebouncer) Init() TextToSpeechDebouncer {
	rv := objc.Send[TextToSpeechDebouncer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechDebouncer) Autorelease() TextToSpeechDebouncer {
	rv := objc.Send[TextToSpeechDebouncer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechDebouncer creates a new TextToSpeechDebouncer instance.
func NewTextToSpeechDebouncer() TextToSpeechDebouncer {
	class := getTextToSpeechDebouncerClass()
	rv := objc.Send[TextToSpeechDebouncer](objc.ID(class.class), objc.Sel("new"))
	return rv
}
