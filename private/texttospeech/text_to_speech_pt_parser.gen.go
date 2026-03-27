// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechPTParser] class.
var (
	_TextToSpeechPTParserClass     TextToSpeechPTParserClass
	_TextToSpeechPTParserClassOnce sync.Once
)

func getTextToSpeechPTParserClass() TextToSpeechPTParserClass {
	_TextToSpeechPTParserClassOnce.Do(func() {
		_TextToSpeechPTParserClass = TextToSpeechPTParserClass{class: objc.GetClass("TextToSpeech.PTParser")}
	})
	return _TextToSpeechPTParserClass
}

// GetTextToSpeechPTParserClass returns the class object for TextToSpeech.PTParser.
func GetTextToSpeechPTParserClass() TextToSpeechPTParserClass {
	return getTextToSpeechPTParserClass()
}

type TextToSpeechPTParserClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechPTParserClass) Alloc() TextToSpeechPTParser {
	rv := objc.Send[TextToSpeechPTParser](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.PTParser
type TextToSpeechPTParser struct {
	objectivec.Object
}

// TextToSpeechPTParserFromID constructs a [TextToSpeechPTParser] from an objc.ID.
func TextToSpeechPTParserFromID(id objc.ID) TextToSpeechPTParser {
	return TextToSpeechPTParser{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechPTParser implements ITextToSpeechPTParser.
var _ ITextToSpeechPTParser = TextToSpeechPTParser{}

// An interface definition for the [TextToSpeechPTParser] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.PTParser
type ITextToSpeechPTParser interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechPTParser) Init() TextToSpeechPTParser {
	rv := objc.Send[TextToSpeechPTParser](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechPTParser) Autorelease() TextToSpeechPTParser {
	rv := objc.Send[TextToSpeechPTParser](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechPTParser creates a new TextToSpeechPTParser instance.
func NewTextToSpeechPTParser() TextToSpeechPTParser {
	class := getTextToSpeechPTParserClass()
	rv := objc.Send[TextToSpeechPTParser](objc.ID(class.class), objc.Sel("new"))
	return rv
}

