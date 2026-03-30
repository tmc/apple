// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechStringStream] class.
var (
	_TextToSpeechStringStreamClass     TextToSpeechStringStreamClass
	_TextToSpeechStringStreamClassOnce sync.Once
)

func getTextToSpeechStringStreamClass() TextToSpeechStringStreamClass {
	_TextToSpeechStringStreamClassOnce.Do(func() {
		_TextToSpeechStringStreamClass = TextToSpeechStringStreamClass{class: objc.GetClass("TextToSpeech.StringStream")}
	})
	return _TextToSpeechStringStreamClass
}

// GetTextToSpeechStringStreamClass returns the class object for TextToSpeech.StringStream.
func GetTextToSpeechStringStreamClass() TextToSpeechStringStreamClass {
	return getTextToSpeechStringStreamClass()
}

type TextToSpeechStringStreamClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechStringStreamClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechStringStreamClass) Alloc() TextToSpeechStringStream {
	rv := objc.Send[TextToSpeechStringStream](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.StringStream
type TextToSpeechStringStream struct {
	objectivec.Object
}

// TextToSpeechStringStreamFromID constructs a [TextToSpeechStringStream] from an objc.ID.
func TextToSpeechStringStreamFromID(id objc.ID) TextToSpeechStringStream {
	return TextToSpeechStringStream{objectivec.Object{ID: id}}
}

// NOTE: TextToSpeechStringStream struct embeds objectivec.Object (parent type unavailable) but
// ITextToSpeechStringStream embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TextToSpeechStringStream] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.StringStream
type ITextToSpeechStringStream interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechStringStream) Init() TextToSpeechStringStream {
	rv := objc.Send[TextToSpeechStringStream](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechStringStream) Autorelease() TextToSpeechStringStream {
	rv := objc.Send[TextToSpeechStringStream](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechStringStream creates a new TextToSpeechStringStream instance.
func NewTextToSpeechStringStream() TextToSpeechStringStream {
	class := getTextToSpeechStringStreamClass()
	rv := objc.Send[TextToSpeechStringStream](objc.ID(class.class), objc.Sel("new"))
	return rv
}
