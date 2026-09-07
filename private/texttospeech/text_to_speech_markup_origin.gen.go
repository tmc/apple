// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechMarkupOrigin] class.
var (
	_TextToSpeechMarkupOriginClass     TextToSpeechMarkupOriginClass
	_TextToSpeechMarkupOriginClassOnce sync.Once
)

func getTextToSpeechMarkupOriginClass() TextToSpeechMarkupOriginClass {
	_TextToSpeechMarkupOriginClassOnce.Do(func() {
		_TextToSpeechMarkupOriginClass = TextToSpeechMarkupOriginClass{class: objc.GetClass("TextToSpeech.MarkupOrigin")}
	})
	return _TextToSpeechMarkupOriginClass
}

// GetTextToSpeechMarkupOriginClass returns the class object for TextToSpeech.MarkupOrigin.
func GetTextToSpeechMarkupOriginClass() TextToSpeechMarkupOriginClass {
	return getTextToSpeechMarkupOriginClass()
}

type TextToSpeechMarkupOriginClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechMarkupOriginClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechMarkupOriginClass) Alloc() TextToSpeechMarkupOrigin {
	rv := objc.SendIfResponds[TextToSpeechMarkupOrigin](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechMarkupOrigin struct {
	objectivec.Object
}

// TextToSpeechMarkupOriginFromID constructs a [TextToSpeechMarkupOrigin] from an objc.ID.
func TextToSpeechMarkupOriginFromID(id objc.ID) TextToSpeechMarkupOrigin {
	return TextToSpeechMarkupOrigin{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechMarkupOrigin implements ITextToSpeechMarkupOrigin.
var _ ITextToSpeechMarkupOrigin = TextToSpeechMarkupOrigin{}

// An interface definition for the [TextToSpeechMarkupOrigin] class.
type ITextToSpeechMarkupOrigin interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechMarkupOrigin) Init() TextToSpeechMarkupOrigin {
	rv := objc.SendIfResponds[TextToSpeechMarkupOrigin](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechMarkupOrigin) Autorelease() TextToSpeechMarkupOrigin {
	rv := objc.SendIfResponds[TextToSpeechMarkupOrigin](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechMarkupOrigin creates a new TextToSpeechMarkupOrigin instance.
func NewTextToSpeechMarkupOrigin() TextToSpeechMarkupOrigin {
	class := getTextToSpeechMarkupOriginClass()
	rv := objc.SendIfResponds[TextToSpeechMarkupOrigin](objc.ID(class.class), objc.Sel("new"))
	return rv
}
