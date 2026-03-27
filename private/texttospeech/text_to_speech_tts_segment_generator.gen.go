// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSSegmentGenerator] class.
var (
	_TextToSpeechTTSSegmentGeneratorClass     TextToSpeechTTSSegmentGeneratorClass
	_TextToSpeechTTSSegmentGeneratorClassOnce sync.Once
)

func getTextToSpeechTTSSegmentGeneratorClass() TextToSpeechTTSSegmentGeneratorClass {
	_TextToSpeechTTSSegmentGeneratorClassOnce.Do(func() {
		_TextToSpeechTTSSegmentGeneratorClass = TextToSpeechTTSSegmentGeneratorClass{class: objc.GetClass("TextToSpeech.TTSSegmentGenerator")}
	})
	return _TextToSpeechTTSSegmentGeneratorClass
}

// GetTextToSpeechTTSSegmentGeneratorClass returns the class object for TextToSpeech.TTSSegmentGenerator.
func GetTextToSpeechTTSSegmentGeneratorClass() TextToSpeechTTSSegmentGeneratorClass {
	return getTextToSpeechTTSSegmentGeneratorClass()
}

type TextToSpeechTTSSegmentGeneratorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSSegmentGeneratorClass) Alloc() TextToSpeechTTSSegmentGenerator {
	rv := objc.Send[TextToSpeechTTSSegmentGenerator](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSegmentGenerator
type TextToSpeechTTSSegmentGenerator struct {
	objectivec.Object
}

// TextToSpeechTTSSegmentGeneratorFromID constructs a [TextToSpeechTTSSegmentGenerator] from an objc.ID.
func TextToSpeechTTSSegmentGeneratorFromID(id objc.ID) TextToSpeechTTSSegmentGenerator {
	return TextToSpeechTTSSegmentGenerator{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechTTSSegmentGenerator implements ITextToSpeechTTSSegmentGenerator.
var _ ITextToSpeechTTSSegmentGenerator = TextToSpeechTTSSegmentGenerator{}

// An interface definition for the [TextToSpeechTTSSegmentGenerator] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSegmentGenerator
type ITextToSpeechTTSSegmentGenerator interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSSegmentGenerator) Init() TextToSpeechTTSSegmentGenerator {
	rv := objc.Send[TextToSpeechTTSSegmentGenerator](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSSegmentGenerator) Autorelease() TextToSpeechTTSSegmentGenerator {
	rv := objc.Send[TextToSpeechTTSSegmentGenerator](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSSegmentGenerator creates a new TextToSpeechTTSSegmentGenerator instance.
func NewTextToSpeechTTSSegmentGenerator() TextToSpeechTTSSegmentGenerator {
	class := getTextToSpeechTTSSegmentGeneratorClass()
	rv := objc.Send[TextToSpeechTTSSegmentGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

