// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/avfaudio"
)

// The class instance for the [TextToSpeechTTSAURenderer] class.
var (
	_TextToSpeechTTSAURendererClass     TextToSpeechTTSAURendererClass
	_TextToSpeechTTSAURendererClassOnce sync.Once
)

func getTextToSpeechTTSAURendererClass() TextToSpeechTTSAURendererClass {
	_TextToSpeechTTSAURendererClassOnce.Do(func() {
		_TextToSpeechTTSAURendererClass = TextToSpeechTTSAURendererClass{class: objc.GetClass("TextToSpeech.TTSAURenderer")}
	})
	return _TextToSpeechTTSAURendererClass
}

// GetTextToSpeechTTSAURendererClass returns the class object for TextToSpeech.TTSAURenderer.
func GetTextToSpeechTTSAURendererClass() TextToSpeechTTSAURendererClass {
	return getTextToSpeechTTSAURendererClass()
}

type TextToSpeechTTSAURendererClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSAURendererClass) Alloc() TextToSpeechTTSAURenderer {
	rv := objc.Send[TextToSpeechTTSAURenderer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSAURenderer
type TextToSpeechTTSAURenderer struct {
	SwiftNativeNSObject
}

// TextToSpeechTTSAURendererFromID constructs a [TextToSpeechTTSAURenderer] from an objc.ID.
func TextToSpeechTTSAURendererFromID(id objc.ID) TextToSpeechTTSAURenderer {
	return TextToSpeechTTSAURenderer{SwiftNativeNSObject: SwiftNativeNSObjectFromID(id)}
}
// Ensure TextToSpeechTTSAURenderer implements ITextToSpeechTTSAURenderer.
var _ ITextToSpeechTTSAURenderer = TextToSpeechTTSAURenderer{}

// An interface definition for the [TextToSpeechTTSAURenderer] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSAURenderer
type ITextToSpeechTTSAURenderer interface {
	ISwiftNativeNSObject
}

// Init initializes the instance.
func (t TextToSpeechTTSAURenderer) Init() TextToSpeechTTSAURenderer {
	rv := objc.Send[TextToSpeechTTSAURenderer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSAURenderer) Autorelease() TextToSpeechTTSAURenderer {
	rv := objc.Send[TextToSpeechTTSAURenderer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSAURenderer creates a new TextToSpeechTTSAURenderer instance.
func NewTextToSpeechTTSAURenderer() TextToSpeechTTSAURenderer {
	class := getTextToSpeechTTSAURendererClass()
	rv := objc.Send[TextToSpeechTTSAURenderer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSAURenderer/formatForVoice:completionHandler:
func (_TextToSpeechTTSAURendererClass TextToSpeechTTSAURendererClass) FormatForVoiceCompletionHandler(voice avfaudio.AVSpeechSynthesisProviderVoice, handler BoolErrorHandler) {
_block1, _cleanup1 := NewBoolErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](objc.ID(_TextToSpeechTTSAURendererClass.class), objc.Sel("formatForVoice:completionHandler:"), voice, _block1)
}

// FormatForVoice is a synchronous wrapper around [TextToSpeechTTSAURenderer.FormatForVoiceCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (tc TextToSpeechTTSAURendererClass) FormatForVoice(ctx context.Context, voice avfaudio.AVSpeechSynthesisProviderVoice) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	tc.FormatForVoiceCompletionHandler(voice, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

