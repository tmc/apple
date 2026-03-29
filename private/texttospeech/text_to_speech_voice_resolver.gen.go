// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechVoiceResolver] class.
var (
	_TextToSpeechVoiceResolverClass     TextToSpeechVoiceResolverClass
	_TextToSpeechVoiceResolverClassOnce sync.Once
)

func getTextToSpeechVoiceResolverClass() TextToSpeechVoiceResolverClass {
	_TextToSpeechVoiceResolverClassOnce.Do(func() {
		_TextToSpeechVoiceResolverClass = TextToSpeechVoiceResolverClass{class: objc.GetClass("TextToSpeech.VoiceResolver")}
	})
	return _TextToSpeechVoiceResolverClass
}

// GetTextToSpeechVoiceResolverClass returns the class object for TextToSpeech.VoiceResolver.
func GetTextToSpeechVoiceResolverClass() TextToSpeechVoiceResolverClass {
	return getTextToSpeechVoiceResolverClass()
}

type TextToSpeechVoiceResolverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechVoiceResolverClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechVoiceResolverClass) Alloc() TextToSpeechVoiceResolver {
	rv := objc.Send[TextToSpeechVoiceResolver](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TextToSpeechVoiceResolver.CurrentLocaleIdentifiersWithCompletionHandler]
//   - [TextToSpeechVoiceResolver.CurrentSystemLocaleIdentifierWithCompletionHandler]
//   - [TextToSpeechVoiceResolver.CurrentSystemLocaleWithCompletionHandler]
//   - [TextToSpeechVoiceResolver.FallbackForVoiceCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForIdentifierCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForIdentifierPreferringLanguageCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForLocaleCompletionHandler]
//   - [TextToSpeechVoiceResolver.VoiceForLocaleIdentifierCompletionHandler]
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver
type TextToSpeechVoiceResolver struct {
	objectivec.Object
}

// TextToSpeechVoiceResolverFromID constructs a [TextToSpeechVoiceResolver] from an objc.ID.
func TextToSpeechVoiceResolverFromID(id objc.ID) TextToSpeechVoiceResolver {
	return TextToSpeechVoiceResolver{objectivec.Object{ID: id}}
}
// NOTE: TextToSpeechVoiceResolver struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechVoiceResolver embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechVoiceResolver] class.
//
// # Methods
//
//   - [ITextToSpeechVoiceResolver.CurrentLocaleIdentifiersWithCompletionHandler]
//   - [ITextToSpeechVoiceResolver.CurrentSystemLocaleIdentifierWithCompletionHandler]
//   - [ITextToSpeechVoiceResolver.CurrentSystemLocaleWithCompletionHandler]
//   - [ITextToSpeechVoiceResolver.FallbackForVoiceCompletionHandler]
//   - [ITextToSpeechVoiceResolver.VoiceForIdentifierCompletionHandler]
//   - [ITextToSpeechVoiceResolver.VoiceForIdentifierPreferringLanguageCompletionHandler]
//   - [ITextToSpeechVoiceResolver.VoiceForLocaleCompletionHandler]
//   - [ITextToSpeechVoiceResolver.VoiceForLocaleIdentifierCompletionHandler]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver
type ITextToSpeechVoiceResolver interface {
	ISwiftNativeNSObject

	// Topic: Methods

	CurrentLocaleIdentifiersWithCompletionHandler(handler ErrorHandler)
	CurrentSystemLocaleIdentifierWithCompletionHandler(handler ErrorHandler)
	CurrentSystemLocaleWithCompletionHandler(handler ErrorHandler)
	FallbackForVoiceCompletionHandler(voice ITTSSpeechVoice, handler ErrorHandler)
	VoiceForIdentifierCompletionHandler(identifier string, handler ErrorHandler)
	VoiceForIdentifierPreferringLanguageCompletionHandler(identifier objectivec.IObject, language objectivec.IObject, handler ErrorHandler)
	VoiceForLocaleCompletionHandler(locale foundation.NSLocale, handler ErrorHandler)
	VoiceForLocaleIdentifierCompletionHandler(identifier objectivec.IObject, handler ErrorHandler)
}

// Init initializes the instance.
func (t TextToSpeechVoiceResolver) Init() TextToSpeechVoiceResolver {
	rv := objc.Send[TextToSpeechVoiceResolver](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechVoiceResolver) Autorelease() TextToSpeechVoiceResolver {
	rv := objc.Send[TextToSpeechVoiceResolver](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechVoiceResolver creates a new TextToSpeechVoiceResolver instance.
func NewTextToSpeechVoiceResolver() TextToSpeechVoiceResolver {
	class := getTextToSpeechVoiceResolverClass()
	rv := objc.Send[TextToSpeechVoiceResolver](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver/currentLocaleIdentifiersWithCompletionHandler:
func (t TextToSpeechVoiceResolver) CurrentLocaleIdentifiersWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("currentLocaleIdentifiersWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver/currentSystemLocaleIdentifierWithCompletionHandler:
func (t TextToSpeechVoiceResolver) CurrentSystemLocaleIdentifierWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("currentSystemLocaleIdentifierWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver/currentSystemLocaleWithCompletionHandler:
func (t TextToSpeechVoiceResolver) CurrentSystemLocaleWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("currentSystemLocaleWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver/fallbackForVoice:completionHandler:
func (t TextToSpeechVoiceResolver) FallbackForVoiceCompletionHandler(voice ITTSSpeechVoice, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("fallbackForVoice:completionHandler:"), voice, _block1)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver/voiceForIdentifier:completionHandler:
func (t TextToSpeechVoiceResolver) VoiceForIdentifierCompletionHandler(identifier string, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("voiceForIdentifier:completionHandler:"), objc.String(identifier), _block1)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver/voiceForIdentifier:preferringLanguage:completionHandler:
func (t TextToSpeechVoiceResolver) VoiceForIdentifierPreferringLanguageCompletionHandler(identifier objectivec.IObject, language objectivec.IObject, handler ErrorHandler) {
_block2, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("voiceForIdentifier:preferringLanguage:completionHandler:"), identifier, language, _block2)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver/voiceForLocale:completionHandler:
func (t TextToSpeechVoiceResolver) VoiceForLocaleCompletionHandler(locale foundation.NSLocale, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("voiceForLocale:completionHandler:"), locale, _block1)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver/voiceForLocaleIdentifier:completionHandler:
func (t TextToSpeechVoiceResolver) VoiceForLocaleIdentifierCompletionHandler(identifier objectivec.IObject, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("voiceForLocaleIdentifier:completionHandler:"), identifier, _block1)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.VoiceResolver/setShared:
func (_TextToSpeechVoiceResolverClass TextToSpeechVoiceResolverClass) SetShared(shared objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_TextToSpeechVoiceResolverClass.class), objc.Sel("setShared:"), shared)
}

// CurrentLocaleIdentifiers is a synchronous wrapper around [TextToSpeechVoiceResolver.CurrentLocaleIdentifiersWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechVoiceResolver) CurrentLocaleIdentifiers(ctx context.Context) error {
	done := make(chan error, 1)
	t.CurrentLocaleIdentifiersWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CurrentSystemLocaleIdentifier is a synchronous wrapper around [TextToSpeechVoiceResolver.CurrentSystemLocaleIdentifierWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechVoiceResolver) CurrentSystemLocaleIdentifier(ctx context.Context) error {
	done := make(chan error, 1)
	t.CurrentSystemLocaleIdentifierWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CurrentSystemLocale is a synchronous wrapper around [TextToSpeechVoiceResolver.CurrentSystemLocaleWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechVoiceResolver) CurrentSystemLocale(ctx context.Context) error {
	done := make(chan error, 1)
	t.CurrentSystemLocaleWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// FallbackForVoice is a synchronous wrapper around [TextToSpeechVoiceResolver.FallbackForVoiceCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechVoiceResolver) FallbackForVoice(ctx context.Context, voice ITTSSpeechVoice) error {
	done := make(chan error, 1)
	t.FallbackForVoiceCompletionHandler(voice, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// VoiceForIdentifier is a synchronous wrapper around [TextToSpeechVoiceResolver.VoiceForIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechVoiceResolver) VoiceForIdentifier(ctx context.Context, identifier string) error {
	done := make(chan error, 1)
	t.VoiceForIdentifierCompletionHandler(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// VoiceForIdentifierPreferringLanguage is a synchronous wrapper around [TextToSpeechVoiceResolver.VoiceForIdentifierPreferringLanguageCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechVoiceResolver) VoiceForIdentifierPreferringLanguage(ctx context.Context, identifier objectivec.IObject, language objectivec.IObject) error {
	done := make(chan error, 1)
	t.VoiceForIdentifierPreferringLanguageCompletionHandler(identifier, language, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// VoiceForLocale is a synchronous wrapper around [TextToSpeechVoiceResolver.VoiceForLocaleCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechVoiceResolver) VoiceForLocale(ctx context.Context, locale foundation.NSLocale) error {
	done := make(chan error, 1)
	t.VoiceForLocaleCompletionHandler(locale, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// VoiceForLocaleIdentifier is a synchronous wrapper around [TextToSpeechVoiceResolver.VoiceForLocaleIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechVoiceResolver) VoiceForLocaleIdentifier(ctx context.Context, identifier objectivec.IObject) error {
	done := make(chan error, 1)
	t.VoiceForLocaleIdentifierCompletionHandler(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

