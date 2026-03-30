// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechCoreSynthesisVoiceShim] class.
var (
	_TextToSpeechCoreSynthesisVoiceShimClass     TextToSpeechCoreSynthesisVoiceShimClass
	_TextToSpeechCoreSynthesisVoiceShimClassOnce sync.Once
)

func getTextToSpeechCoreSynthesisVoiceShimClass() TextToSpeechCoreSynthesisVoiceShimClass {
	_TextToSpeechCoreSynthesisVoiceShimClassOnce.Do(func() {
		_TextToSpeechCoreSynthesisVoiceShimClass = TextToSpeechCoreSynthesisVoiceShimClass{class: objc.GetClass("TextToSpeech.CoreSynthesisVoiceShim")}
	})
	return _TextToSpeechCoreSynthesisVoiceShimClass
}

// GetTextToSpeechCoreSynthesisVoiceShimClass returns the class object for TextToSpeech.CoreSynthesisVoiceShim.
func GetTextToSpeechCoreSynthesisVoiceShimClass() TextToSpeechCoreSynthesisVoiceShimClass {
	return getTextToSpeechCoreSynthesisVoiceShimClass()
}

type TextToSpeechCoreSynthesisVoiceShimClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechCoreSynthesisVoiceShimClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechCoreSynthesisVoiceShimClass) Alloc() TextToSpeechCoreSynthesisVoiceShim {
	rv := objc.Send[TextToSpeechCoreSynthesisVoiceShim](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoiceWithIdentifier]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoicesIncludingSiri]
//   - [TextToSpeechCoreSynthesisVoiceShim.InternalVoicesIncludingSiriCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.PublicVoices]
//   - [TextToSpeechCoreSynthesisVoiceShim.PublicVoicesWithCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ReloadPublicResolver]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoiceWithIdentifier]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoicesWithOnlyInstalled]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourceVoicesWithOnlyInstalledCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourcesWithLanguageCode]
//   - [TextToSpeechCoreSynthesisVoiceShim.ResourcesWithLanguageCodeCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithIdentifier]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithLanguageCode]
//   - [TextToSpeechCoreSynthesisVoiceShim.VoiceWithLanguageCodeCompletionHandler]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim
type TextToSpeechCoreSynthesisVoiceShim struct {
	objectivec.Object
}

// TextToSpeechCoreSynthesisVoiceShimFromID constructs a [TextToSpeechCoreSynthesisVoiceShim] from an objc.ID.
func TextToSpeechCoreSynthesisVoiceShimFromID(id objc.ID) TextToSpeechCoreSynthesisVoiceShim {
	return TextToSpeechCoreSynthesisVoiceShim{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechCoreSynthesisVoiceShim implements ITextToSpeechCoreSynthesisVoiceShim.
var _ ITextToSpeechCoreSynthesisVoiceShim = TextToSpeechCoreSynthesisVoiceShim{}

// An interface definition for the [TextToSpeechCoreSynthesisVoiceShim] class.
//
// # Methods
//
//   - [ITextToSpeechCoreSynthesisVoiceShim.InternalVoiceWithIdentifier]
//   - [ITextToSpeechCoreSynthesisVoiceShim.InternalVoiceWithIdentifierCompletionHandler]
//   - [ITextToSpeechCoreSynthesisVoiceShim.InternalVoicesIncludingSiri]
//   - [ITextToSpeechCoreSynthesisVoiceShim.InternalVoicesIncludingSiriCompletionHandler]
//   - [ITextToSpeechCoreSynthesisVoiceShim.PublicVoices]
//   - [ITextToSpeechCoreSynthesisVoiceShim.PublicVoicesWithCompletionHandler]
//   - [ITextToSpeechCoreSynthesisVoiceShim.ReloadPublicResolver]
//   - [ITextToSpeechCoreSynthesisVoiceShim.ResourceVoiceWithIdentifier]
//   - [ITextToSpeechCoreSynthesisVoiceShim.ResourceVoiceWithIdentifierCompletionHandler]
//   - [ITextToSpeechCoreSynthesisVoiceShim.ResourceVoicesWithOnlyInstalled]
//   - [ITextToSpeechCoreSynthesisVoiceShim.ResourceVoicesWithOnlyInstalledCompletionHandler]
//   - [ITextToSpeechCoreSynthesisVoiceShim.ResourcesWithLanguageCode]
//   - [ITextToSpeechCoreSynthesisVoiceShim.ResourcesWithLanguageCodeCompletionHandler]
//   - [ITextToSpeechCoreSynthesisVoiceShim.VoiceWithIdentifier]
//   - [ITextToSpeechCoreSynthesisVoiceShim.VoiceWithIdentifierCompletionHandler]
//   - [ITextToSpeechCoreSynthesisVoiceShim.VoiceWithLanguageCode]
//   - [ITextToSpeechCoreSynthesisVoiceShim.VoiceWithLanguageCodeCompletionHandler]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim
type ITextToSpeechCoreSynthesisVoiceShim interface {
	objectivec.IObject

	// Topic: Methods

	InternalVoiceWithIdentifier(identifier objectivec.IObject) objectivec.IObject
	InternalVoiceWithIdentifierCompletionHandler(identifier string, handler ErrorHandler)
	InternalVoicesIncludingSiri(siri bool) objectivec.IObject
	InternalVoicesIncludingSiriCompletionHandler(siri bool, handler ErrorHandler)
	PublicVoices() objectivec.IObject
	PublicVoicesWithCompletionHandler(handler ErrorHandler)
	ReloadPublicResolver()
	ResourceVoiceWithIdentifier(identifier objectivec.IObject) objectivec.IObject
	ResourceVoiceWithIdentifierCompletionHandler(identifier string, handler ErrorHandler)
	ResourceVoicesWithOnlyInstalled(installed bool) objectivec.IObject
	ResourceVoicesWithOnlyInstalledCompletionHandler(installed bool, handler ErrorHandler)
	ResourcesWithLanguageCode(code objectivec.IObject) objectivec.IObject
	ResourcesWithLanguageCodeCompletionHandler(code string, handler ErrorHandler)
	VoiceWithIdentifier(identifier objectivec.IObject) objectivec.IObject
	VoiceWithIdentifierCompletionHandler(identifier string, handler ErrorHandler)
	VoiceWithLanguageCode(code objectivec.IObject) objectivec.IObject
	VoiceWithLanguageCodeCompletionHandler(code string, handler ErrorHandler)
}

// Init initializes the instance.
func (t TextToSpeechCoreSynthesisVoiceShim) Init() TextToSpeechCoreSynthesisVoiceShim {
	rv := objc.Send[TextToSpeechCoreSynthesisVoiceShim](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechCoreSynthesisVoiceShim) Autorelease() TextToSpeechCoreSynthesisVoiceShim {
	rv := objc.Send[TextToSpeechCoreSynthesisVoiceShim](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechCoreSynthesisVoiceShim creates a new TextToSpeechCoreSynthesisVoiceShim instance.
func NewTextToSpeechCoreSynthesisVoiceShim() TextToSpeechCoreSynthesisVoiceShim {
	class := getTextToSpeechCoreSynthesisVoiceShimClass()
	rv := objc.Send[TextToSpeechCoreSynthesisVoiceShim](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/internalVoiceWithIdentifier:
func (t TextToSpeechCoreSynthesisVoiceShim) InternalVoiceWithIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("internalVoiceWithIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/internalVoiceWithIdentifier:completionHandler:
func (t TextToSpeechCoreSynthesisVoiceShim) InternalVoiceWithIdentifierCompletionHandler(identifier string, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("internalVoiceWithIdentifier:completionHandler:"), objc.String(identifier), _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/internalVoicesIncludingSiri:
func (t TextToSpeechCoreSynthesisVoiceShim) InternalVoicesIncludingSiri(siri bool) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("internalVoicesIncludingSiri:"), siri)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/internalVoicesIncludingSiri:completionHandler:
func (t TextToSpeechCoreSynthesisVoiceShim) InternalVoicesIncludingSiriCompletionHandler(siri bool, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("internalVoicesIncludingSiri:completionHandler:"), siri, _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/publicVoices
func (t TextToSpeechCoreSynthesisVoiceShim) PublicVoices() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("publicVoices"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/publicVoicesWithCompletionHandler:
func (t TextToSpeechCoreSynthesisVoiceShim) PublicVoicesWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("publicVoicesWithCompletionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/reloadPublicResolver
func (t TextToSpeechCoreSynthesisVoiceShim) ReloadPublicResolver() {
	objc.Send[objc.ID](t.ID, objc.Sel("reloadPublicResolver"))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/resourceVoiceWithIdentifier:
func (t TextToSpeechCoreSynthesisVoiceShim) ResourceVoiceWithIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resourceVoiceWithIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/resourceVoiceWithIdentifier:completionHandler:
func (t TextToSpeechCoreSynthesisVoiceShim) ResourceVoiceWithIdentifierCompletionHandler(identifier string, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("resourceVoiceWithIdentifier:completionHandler:"), objc.String(identifier), _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/resourceVoicesWithOnlyInstalled:
func (t TextToSpeechCoreSynthesisVoiceShim) ResourceVoicesWithOnlyInstalled(installed bool) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resourceVoicesWithOnlyInstalled:"), installed)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/resourceVoicesWithOnlyInstalled:completionHandler:
func (t TextToSpeechCoreSynthesisVoiceShim) ResourceVoicesWithOnlyInstalledCompletionHandler(installed bool, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("resourceVoicesWithOnlyInstalled:completionHandler:"), installed, _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/resourcesWithLanguageCode:
func (t TextToSpeechCoreSynthesisVoiceShim) ResourcesWithLanguageCode(code objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("resourcesWithLanguageCode:"), code)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/resourcesWithLanguageCode:completionHandler:
func (t TextToSpeechCoreSynthesisVoiceShim) ResourcesWithLanguageCodeCompletionHandler(code string, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("resourcesWithLanguageCode:completionHandler:"), objc.String(code), _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/voiceWithIdentifier:
func (t TextToSpeechCoreSynthesisVoiceShim) VoiceWithIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceWithIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/voiceWithIdentifier:completionHandler:
func (t TextToSpeechCoreSynthesisVoiceShim) VoiceWithIdentifierCompletionHandler(identifier string, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("voiceWithIdentifier:completionHandler:"), objc.String(identifier), _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/voiceWithLanguageCode:
func (t TextToSpeechCoreSynthesisVoiceShim) VoiceWithLanguageCode(code objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceWithLanguageCode:"), code)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/voiceWithLanguageCode:completionHandler:
func (t TextToSpeechCoreSynthesisVoiceShim) VoiceWithLanguageCodeCompletionHandler(code string, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("voiceWithLanguageCode:completionHandler:"), objc.String(code), _block1)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesisVoiceShim/setShared:
func (_TextToSpeechCoreSynthesisVoiceShimClass TextToSpeechCoreSynthesisVoiceShimClass) SetShared(shared objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_TextToSpeechCoreSynthesisVoiceShimClass.class), objc.Sel("setShared:"), shared)
}

// InternalVoiceWithIdentifierSync is a synchronous wrapper around [TextToSpeechCoreSynthesisVoiceShim.InternalVoiceWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesisVoiceShim) InternalVoiceWithIdentifierSync(ctx context.Context, identifier string) error {
	done := make(chan error, 1)
	t.InternalVoiceWithIdentifierCompletionHandler(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InternalVoicesIncludingSiriSync is a synchronous wrapper around [TextToSpeechCoreSynthesisVoiceShim.InternalVoicesIncludingSiriCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesisVoiceShim) InternalVoicesIncludingSiriSync(ctx context.Context, siri bool) error {
	done := make(chan error, 1)
	t.InternalVoicesIncludingSiriCompletionHandler(siri, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PublicVoicesSync is a synchronous wrapper around [TextToSpeechCoreSynthesisVoiceShim.PublicVoicesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesisVoiceShim) PublicVoicesSync(ctx context.Context) error {
	done := make(chan error, 1)
	t.PublicVoicesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ResourceVoiceWithIdentifierSync is a synchronous wrapper around [TextToSpeechCoreSynthesisVoiceShim.ResourceVoiceWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesisVoiceShim) ResourceVoiceWithIdentifierSync(ctx context.Context, identifier string) error {
	done := make(chan error, 1)
	t.ResourceVoiceWithIdentifierCompletionHandler(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ResourceVoicesWithOnlyInstalledSync is a synchronous wrapper around [TextToSpeechCoreSynthesisVoiceShim.ResourceVoicesWithOnlyInstalledCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesisVoiceShim) ResourceVoicesWithOnlyInstalledSync(ctx context.Context, installed bool) error {
	done := make(chan error, 1)
	t.ResourceVoicesWithOnlyInstalledCompletionHandler(installed, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ResourcesWithLanguageCodeSync is a synchronous wrapper around [TextToSpeechCoreSynthesisVoiceShim.ResourcesWithLanguageCodeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesisVoiceShim) ResourcesWithLanguageCodeSync(ctx context.Context, code string) error {
	done := make(chan error, 1)
	t.ResourcesWithLanguageCodeCompletionHandler(code, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// VoiceWithIdentifierSync is a synchronous wrapper around [TextToSpeechCoreSynthesisVoiceShim.VoiceWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesisVoiceShim) VoiceWithIdentifierSync(ctx context.Context, identifier string) error {
	done := make(chan error, 1)
	t.VoiceWithIdentifierCompletionHandler(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// VoiceWithLanguageCodeSync is a synchronous wrapper around [TextToSpeechCoreSynthesisVoiceShim.VoiceWithLanguageCodeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesisVoiceShim) VoiceWithLanguageCodeSync(ctx context.Context, code string) error {
	done := make(chan error, 1)
	t.VoiceWithLanguageCodeCompletionHandler(code, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
