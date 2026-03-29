// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechSpokenContentSelectionProvider] class.
var (
	_TextToSpeechSpokenContentSelectionProviderClass     TextToSpeechSpokenContentSelectionProviderClass
	_TextToSpeechSpokenContentSelectionProviderClassOnce sync.Once
)

func getTextToSpeechSpokenContentSelectionProviderClass() TextToSpeechSpokenContentSelectionProviderClass {
	_TextToSpeechSpokenContentSelectionProviderClassOnce.Do(func() {
		_TextToSpeechSpokenContentSelectionProviderClass = TextToSpeechSpokenContentSelectionProviderClass{class: objc.GetClass("TextToSpeech.SpokenContentSelectionProvider")}
	})
	return _TextToSpeechSpokenContentSelectionProviderClass
}

// GetTextToSpeechSpokenContentSelectionProviderClass returns the class object for TextToSpeech.SpokenContentSelectionProvider.
func GetTextToSpeechSpokenContentSelectionProviderClass() TextToSpeechSpokenContentSelectionProviderClass {
	return getTextToSpeechSpokenContentSelectionProviderClass()
}

type TextToSpeechSpokenContentSelectionProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechSpokenContentSelectionProviderClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechSpokenContentSelectionProviderClass) Alloc() TextToSpeechSpokenContentSelectionProvider {
	rv := objc.Send[TextToSpeechSpokenContentSelectionProvider](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TextToSpeechSpokenContentSelectionProvider.SiriVoiceChanged]
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SpokenContentSelectionProvider
type TextToSpeechSpokenContentSelectionProvider struct {
	objectivec.Object
}

// TextToSpeechSpokenContentSelectionProviderFromID constructs a [TextToSpeechSpokenContentSelectionProvider] from an objc.ID.
func TextToSpeechSpokenContentSelectionProviderFromID(id objc.ID) TextToSpeechSpokenContentSelectionProvider {
	return TextToSpeechSpokenContentSelectionProvider{objectivec.Object{ID: id}}
}
// NOTE: TextToSpeechSpokenContentSelectionProvider struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechSpokenContentSelectionProvider embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechSpokenContentSelectionProvider] class.
//
// # Methods
//
//   - [ITextToSpeechSpokenContentSelectionProvider.SiriVoiceChanged]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SpokenContentSelectionProvider
type ITextToSpeechSpokenContentSelectionProvider interface {
	objectivec.IObject

	// Topic: Methods

	SiriVoiceChanged()
}

// Init initializes the instance.
func (t TextToSpeechSpokenContentSelectionProvider) Init() TextToSpeechSpokenContentSelectionProvider {
	rv := objc.Send[TextToSpeechSpokenContentSelectionProvider](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechSpokenContentSelectionProvider) Autorelease() TextToSpeechSpokenContentSelectionProvider {
	rv := objc.Send[TextToSpeechSpokenContentSelectionProvider](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechSpokenContentSelectionProvider creates a new TextToSpeechSpokenContentSelectionProvider instance.
func NewTextToSpeechSpokenContentSelectionProvider() TextToSpeechSpokenContentSelectionProvider {
	class := getTextToSpeechSpokenContentSelectionProviderClass()
	rv := objc.Send[TextToSpeechSpokenContentSelectionProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.SpokenContentSelectionProvider/siriVoiceChanged
func (t TextToSpeechSpokenContentSelectionProvider) SiriVoiceChanged() {
	objc.Send[objc.ID](t.ID, objc.Sel("siriVoiceChanged"))
}

