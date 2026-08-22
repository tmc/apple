// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechSSMLParser] class.
var (
	_TextToSpeechSSMLParserClass     TextToSpeechSSMLParserClass
	_TextToSpeechSSMLParserClassOnce sync.Once
)

func getTextToSpeechSSMLParserClass() TextToSpeechSSMLParserClass {
	_TextToSpeechSSMLParserClassOnce.Do(func() {
		_TextToSpeechSSMLParserClass = TextToSpeechSSMLParserClass{class: objc.GetClass("TextToSpeech.SSMLParser")}
	})
	return _TextToSpeechSSMLParserClass
}

// GetTextToSpeechSSMLParserClass returns the class object for TextToSpeech.SSMLParser.
func GetTextToSpeechSSMLParserClass() TextToSpeechSSMLParserClass {
	return getTextToSpeechSSMLParserClass()
}

type TextToSpeechSSMLParserClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechSSMLParserClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechSSMLParserClass) Alloc() TextToSpeechSSMLParser {
	rv := objc.SendIfResponds[TextToSpeechSSMLParser](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

type TextToSpeechSSMLParser struct {
	objectivec.Object
}

// TextToSpeechSSMLParserFromID constructs a [TextToSpeechSSMLParser] from an objc.ID.
func TextToSpeechSSMLParserFromID(id objc.ID) TextToSpeechSSMLParser {
	return TextToSpeechSSMLParser{objectivec.Object{ID: id}}
}

// Ensure TextToSpeechSSMLParser implements ITextToSpeechSSMLParser.
var _ ITextToSpeechSSMLParser = TextToSpeechSSMLParser{}

// An interface definition for the [TextToSpeechSSMLParser] class.
type ITextToSpeechSSMLParser interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechSSMLParser) Init() TextToSpeechSSMLParser {
	rv := objc.SendIfResponds[TextToSpeechSSMLParser](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechSSMLParser) Autorelease() TextToSpeechSSMLParser {
	rv := objc.SendIfResponds[TextToSpeechSSMLParser](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechSSMLParser creates a new TextToSpeechSSMLParser instance.
func NewTextToSpeechSSMLParser() TextToSpeechSSMLParser {
	class := getTextToSpeechSSMLParserClass()
	rv := objc.SendIfResponds[TextToSpeechSSMLParser](objc.ID(class.class), objc.Sel("new"))
	return rv
}
