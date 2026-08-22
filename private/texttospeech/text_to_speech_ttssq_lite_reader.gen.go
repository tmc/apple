// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSSQLiteReader] class.
var (
	_TextToSpeechTTSSQLiteReaderClass     TextToSpeechTTSSQLiteReaderClass
	_TextToSpeechTTSSQLiteReaderClassOnce sync.Once
)

func getTextToSpeechTTSSQLiteReaderClass() TextToSpeechTTSSQLiteReaderClass {
	_TextToSpeechTTSSQLiteReaderClassOnce.Do(func() {
		_TextToSpeechTTSSQLiteReaderClass = TextToSpeechTTSSQLiteReaderClass{class: objc.GetClass("TextToSpeech.TTSSQLiteReader")}
	})
	return _TextToSpeechTTSSQLiteReaderClass
}

// GetTextToSpeechTTSSQLiteReaderClass returns the class object for TextToSpeech.TTSSQLiteReader.
func GetTextToSpeechTTSSQLiteReaderClass() TextToSpeechTTSSQLiteReaderClass {
	return getTextToSpeechTTSSQLiteReaderClass()
}

type TextToSpeechTTSSQLiteReaderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechTTSSQLiteReaderClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSSQLiteReaderClass) Alloc() TextToSpeechTTSSQLiteReader {
	rv := objc.SendIfResponds[TextToSpeechTTSSQLiteReader](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechTTSSQLiteReader struct {
	objectivec.Object
}

// TextToSpeechTTSSQLiteReaderFromID constructs a [TextToSpeechTTSSQLiteReader] from an objc.ID.
func TextToSpeechTTSSQLiteReaderFromID(id objc.ID) TextToSpeechTTSSQLiteReader {
	return TextToSpeechTTSSQLiteReader{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechTTSSQLiteReader implements ITextToSpeechTTSSQLiteReader.
var _ ITextToSpeechTTSSQLiteReader = TextToSpeechTTSSQLiteReader{}

// An interface definition for the [TextToSpeechTTSSQLiteReader] class.
type ITextToSpeechTTSSQLiteReader interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSSQLiteReader) Init() TextToSpeechTTSSQLiteReader {
	rv := objc.SendIfResponds[TextToSpeechTTSSQLiteReader](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSSQLiteReader) Autorelease() TextToSpeechTTSSQLiteReader {
	rv := objc.SendIfResponds[TextToSpeechTTSSQLiteReader](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSSQLiteReader creates a new TextToSpeechTTSSQLiteReader instance.
func NewTextToSpeechTTSSQLiteReader() TextToSpeechTTSSQLiteReader {
	class := getTextToSpeechTTSSQLiteReaderClass()
	rv := objc.SendIfResponds[TextToSpeechTTSSQLiteReader](objc.ID(class.class), objc.Sel("new"))
	return rv
}
