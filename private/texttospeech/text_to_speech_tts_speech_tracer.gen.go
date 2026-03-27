// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechTTSSpeechTracer] class.
var (
	_TextToSpeechTTSSpeechTracerClass     TextToSpeechTTSSpeechTracerClass
	_TextToSpeechTTSSpeechTracerClassOnce sync.Once
)

func getTextToSpeechTTSSpeechTracerClass() TextToSpeechTTSSpeechTracerClass {
	_TextToSpeechTTSSpeechTracerClassOnce.Do(func() {
		_TextToSpeechTTSSpeechTracerClass = TextToSpeechTTSSpeechTracerClass{class: objc.GetClass("TextToSpeech.TTSSpeechTracer")}
	})
	return _TextToSpeechTTSSpeechTracerClass
}

// GetTextToSpeechTTSSpeechTracerClass returns the class object for TextToSpeech.TTSSpeechTracer.
func GetTextToSpeechTTSSpeechTracerClass() TextToSpeechTTSSpeechTracerClass {
	return getTextToSpeechTTSSpeechTracerClass()
}

type TextToSpeechTTSSpeechTracerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechTTSSpeechTracerClass) Alloc() TextToSpeechTTSSpeechTracer {
	rv := objc.Send[TextToSpeechTTSSpeechTracer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TextToSpeechTTSSpeechTracer.EmitWithEventForIdentifier]
//   - [TextToSpeechTTSSpeechTracer.MakeSpeechJobIdentifier]
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSpeechTracer
type TextToSpeechTTSSpeechTracer struct {
	objectivec.Object
}

// TextToSpeechTTSSpeechTracerFromID constructs a [TextToSpeechTTSSpeechTracer] from an objc.ID.
func TextToSpeechTTSSpeechTracerFromID(id objc.ID) TextToSpeechTTSSpeechTracer {
	return TextToSpeechTTSSpeechTracer{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechTTSSpeechTracer implements ITextToSpeechTTSSpeechTracer.
var _ ITextToSpeechTTSSpeechTracer = TextToSpeechTTSSpeechTracer{}

// An interface definition for the [TextToSpeechTTSSpeechTracer] class.
//
// # Methods
//
//   - [ITextToSpeechTTSSpeechTracer.EmitWithEventForIdentifier]
//   - [ITextToSpeechTTSSpeechTracer.MakeSpeechJobIdentifier]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSpeechTracer
type ITextToSpeechTTSSpeechTracer interface {
	objectivec.IObject

	// Topic: Methods

	EmitWithEventForIdentifier(event objectivec.IObject, identifier objectivec.IObject)
	MakeSpeechJobIdentifier() objectivec.IObject
}

// Init initializes the instance.
func (t TextToSpeechTTSSpeechTracer) Init() TextToSpeechTTSSpeechTracer {
	rv := objc.Send[TextToSpeechTTSSpeechTracer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechTTSSpeechTracer) Autorelease() TextToSpeechTTSSpeechTracer {
	rv := objc.Send[TextToSpeechTTSSpeechTracer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechTTSSpeechTracer creates a new TextToSpeechTTSSpeechTracer instance.
func NewTextToSpeechTTSSpeechTracer() TextToSpeechTTSSpeechTracer {
	class := getTextToSpeechTTSSpeechTracerClass()
	rv := objc.Send[TextToSpeechTTSSpeechTracer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSpeechTracer/emitWithEvent:forIdentifier:
func (t TextToSpeechTTSSpeechTracer) EmitWithEventForIdentifier(event objectivec.IObject, identifier objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("emitWithEvent:forIdentifier:"), event, identifier)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.TTSSpeechTracer/makeSpeechJobIdentifier
func (t TextToSpeechTTSSpeechTracer) MakeSpeechJobIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("makeSpeechJobIdentifier"))
	return objectivec.Object{ID: rv}
}

