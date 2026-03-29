// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechSynthesisMarker] class.
var (
	_AVSpeechSynthesisMarkerClass     AVSpeechSynthesisMarkerClass
	_AVSpeechSynthesisMarkerClassOnce sync.Once
)

func getAVSpeechSynthesisMarkerClass() AVSpeechSynthesisMarkerClass {
	_AVSpeechSynthesisMarkerClassOnce.Do(func() {
		_AVSpeechSynthesisMarkerClass = AVSpeechSynthesisMarkerClass{class: objc.GetClass("AVSpeechSynthesisMarker")}
	})
	return _AVSpeechSynthesisMarkerClass
}

// GetAVSpeechSynthesisMarkerClass returns the class object for AVSpeechSynthesisMarker.
func GetAVSpeechSynthesisMarkerClass() AVSpeechSynthesisMarkerClass {
	return getAVSpeechSynthesisMarkerClass()
}

type AVSpeechSynthesisMarkerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechSynthesisMarkerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechSynthesisMarkerClass) Alloc() AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that contains information about the synthesized audio.
//
// # Creating a marker
//
//   - [AVSpeechSynthesisMarker.InitWithMarkerTypeForTextRangeAtByteSampleOffset]: Creates a marker with a type and location of the request’s text.
//   - [AVSpeechSynthesisMarker.InitWithWordRangeAtByteSampleOffset]: Creates a word marker with a range of the word and offset into the audio buffer.
//   - [AVSpeechSynthesisMarker.InitWithSentenceRangeAtByteSampleOffset]: Creates a sentence marker with a range of the sentence and offset into the audio buffer.
//   - [AVSpeechSynthesisMarker.InitWithParagraphRangeAtByteSampleOffset]: Creates a paragraph marker with a range of the paragraph and offset into the audio buffer.
//   - [AVSpeechSynthesisMarker.InitWithPhonemeStringAtByteSampleOffset]: Creates a phoneme marker with a range of the phoneme and offset into the audio buffer.
//   - [AVSpeechSynthesisMarker.InitWithBookmarkNameAtByteSampleOffset]: Creates a bookmark marker with a name and offset into the audio buffer.
//
// # Inspecting a marker
//
//   - [AVSpeechSynthesisMarker.Mark]: The type that describes the text.
//   - [AVSpeechSynthesisMarker.SetMark]
//   - [AVSpeechSynthesisMarker.BookmarkName]: A string that represents the name of a bookmark.
//   - [AVSpeechSynthesisMarker.SetBookmarkName]
//   - [AVSpeechSynthesisMarker.Phoneme]: A string that represents a distinct sound.
//   - [AVSpeechSynthesisMarker.SetPhoneme]
//   - [AVSpeechSynthesisMarker.TextRange]: The location and length of the request’s text.
//   - [AVSpeechSynthesisMarker.SetTextRange]
//   - [AVSpeechSynthesisMarker.ByteSampleOffset]: The byte offset into the audio buffer.
//   - [AVSpeechSynthesisMarker.SetByteSampleOffset]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker
type AVSpeechSynthesisMarker struct {
	objectivec.Object
}

// AVSpeechSynthesisMarkerFromID constructs a [AVSpeechSynthesisMarker] from an objc.ID.
//
// An object that contains information about the synthesized audio.
func AVSpeechSynthesisMarkerFromID(id objc.ID) AVSpeechSynthesisMarker {
	return AVSpeechSynthesisMarker{objectivec.Object{ID: id}}
}
// NOTE: AVSpeechSynthesisMarker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSpeechSynthesisMarker] class.
//
// # Creating a marker
//
//   - [IAVSpeechSynthesisMarker.InitWithMarkerTypeForTextRangeAtByteSampleOffset]: Creates a marker with a type and location of the request’s text.
//   - [IAVSpeechSynthesisMarker.InitWithWordRangeAtByteSampleOffset]: Creates a word marker with a range of the word and offset into the audio buffer.
//   - [IAVSpeechSynthesisMarker.InitWithSentenceRangeAtByteSampleOffset]: Creates a sentence marker with a range of the sentence and offset into the audio buffer.
//   - [IAVSpeechSynthesisMarker.InitWithParagraphRangeAtByteSampleOffset]: Creates a paragraph marker with a range of the paragraph and offset into the audio buffer.
//   - [IAVSpeechSynthesisMarker.InitWithPhonemeStringAtByteSampleOffset]: Creates a phoneme marker with a range of the phoneme and offset into the audio buffer.
//   - [IAVSpeechSynthesisMarker.InitWithBookmarkNameAtByteSampleOffset]: Creates a bookmark marker with a name and offset into the audio buffer.
//
// # Inspecting a marker
//
//   - [IAVSpeechSynthesisMarker.Mark]: The type that describes the text.
//   - [IAVSpeechSynthesisMarker.SetMark]
//   - [IAVSpeechSynthesisMarker.BookmarkName]: A string that represents the name of a bookmark.
//   - [IAVSpeechSynthesisMarker.SetBookmarkName]
//   - [IAVSpeechSynthesisMarker.Phoneme]: A string that represents a distinct sound.
//   - [IAVSpeechSynthesisMarker.SetPhoneme]
//   - [IAVSpeechSynthesisMarker.TextRange]: The location and length of the request’s text.
//   - [IAVSpeechSynthesisMarker.SetTextRange]
//   - [IAVSpeechSynthesisMarker.ByteSampleOffset]: The byte offset into the audio buffer.
//   - [IAVSpeechSynthesisMarker.SetByteSampleOffset]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker
type IAVSpeechSynthesisMarker interface {
	objectivec.IObject

	// Topic: Creating a marker

	// Creates a marker with a type and location of the request’s text.
	InitWithMarkerTypeForTextRangeAtByteSampleOffset(type_ AVSpeechSynthesisMarkerMark, range_ foundation.NSRange, byteSampleOffset uint) AVSpeechSynthesisMarker
	// Creates a word marker with a range of the word and offset into the audio buffer.
	InitWithWordRangeAtByteSampleOffset(range_ foundation.NSRange, byteSampleOffset int) AVSpeechSynthesisMarker
	// Creates a sentence marker with a range of the sentence and offset into the audio buffer.
	InitWithSentenceRangeAtByteSampleOffset(range_ foundation.NSRange, byteSampleOffset int) AVSpeechSynthesisMarker
	// Creates a paragraph marker with a range of the paragraph and offset into the audio buffer.
	InitWithParagraphRangeAtByteSampleOffset(range_ foundation.NSRange, byteSampleOffset int) AVSpeechSynthesisMarker
	// Creates a phoneme marker with a range of the phoneme and offset into the audio buffer.
	InitWithPhonemeStringAtByteSampleOffset(phoneme string, byteSampleOffset int) AVSpeechSynthesisMarker
	// Creates a bookmark marker with a name and offset into the audio buffer.
	InitWithBookmarkNameAtByteSampleOffset(mark string, byteSampleOffset int) AVSpeechSynthesisMarker

	// Topic: Inspecting a marker

	// The type that describes the text.
	Mark() AVSpeechSynthesisMarkerMark
	SetMark(value AVSpeechSynthesisMarkerMark)
	// A string that represents the name of a bookmark.
	BookmarkName() string
	SetBookmarkName(value string)
	// A string that represents a distinct sound.
	Phoneme() string
	SetPhoneme(value string)
	// The location and length of the request’s text.
	TextRange() foundation.NSRange
	SetTextRange(value foundation.NSRange)
	// The byte offset into the audio buffer.
	ByteSampleOffset() uint
	SetByteSampleOffset(value uint)

	// A block that subclasses use to send marker information to the host.
	SpeechSynthesisOutputMetadataBlock() AVSpeechSynthesisProviderOutputBlock
	SetSpeechSynthesisOutputMetadataBlock(value AVSpeechSynthesisProviderOutputBlock)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s AVSpeechSynthesisMarker) Init() AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechSynthesisMarker) Autorelease() AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesisMarker creates a new AVSpeechSynthesisMarker instance.
func NewAVSpeechSynthesisMarker() AVSpeechSynthesisMarker {
	class := getAVSpeechSynthesisMarkerClass()
	rv := objc.Send[AVSpeechSynthesisMarker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a bookmark marker with a name and offset into the audio buffer.
//
// mark: The name of the bookmark.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(bookmarkName:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithBookmarkNameAtByteSampleOffset(mark string, byteSampleOffset int) AVSpeechSynthesisMarker {
	instance := getAVSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBookmarkName:atByteSampleOffset:"), objc.String(mark), byteSampleOffset)
	return AVSpeechSynthesisMarkerFromID(rv)
}

// Creates a marker with a type and location of the request’s text.
//
// type: The type that describes the text.
//
// range: The location and length of the request’s text.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(markerType:forTextRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithMarkerTypeForTextRangeAtByteSampleOffset(type_ AVSpeechSynthesisMarkerMark, range_ foundation.NSRange, byteSampleOffset uint) AVSpeechSynthesisMarker {
	instance := getAVSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMarkerType:forTextRange:atByteSampleOffset:"), type_, range_, byteSampleOffset)
	return AVSpeechSynthesisMarkerFromID(rv)
}

// Creates a paragraph marker with a range of the paragraph and offset into
// the audio buffer.
//
// range: The location and length of the paragraph.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(paragraphRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithParagraphRangeAtByteSampleOffset(range_ foundation.NSRange, byteSampleOffset int) AVSpeechSynthesisMarker {
	instance := getAVSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParagraphRange:atByteSampleOffset:"), range_, byteSampleOffset)
	return AVSpeechSynthesisMarkerFromID(rv)
}

// Creates a phoneme marker with a range of the phoneme and offset into the
// audio buffer.
//
// phoneme: A string that represents a distinct sound.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(phonemeString:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithPhonemeStringAtByteSampleOffset(phoneme string, byteSampleOffset int) AVSpeechSynthesisMarker {
	instance := getAVSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPhonemeString:atByteSampleOffset:"), objc.String(phoneme), byteSampleOffset)
	return AVSpeechSynthesisMarkerFromID(rv)
}

// Creates a sentence marker with a range of the sentence and offset into the
// audio buffer.
//
// range: The location and length of the word.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(sentenceRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithSentenceRangeAtByteSampleOffset(range_ foundation.NSRange, byteSampleOffset int) AVSpeechSynthesisMarker {
	instance := getAVSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSentenceRange:atByteSampleOffset:"), range_, byteSampleOffset)
	return AVSpeechSynthesisMarkerFromID(rv)
}

// Creates a word marker with a range of the word and offset into the audio
// buffer.
//
// range: The location and length of the word.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(wordRange:atByteSampleOffset:)
func NewSpeechSynthesisMarkerWithWordRangeAtByteSampleOffset(range_ foundation.NSRange, byteSampleOffset int) AVSpeechSynthesisMarker {
	instance := getAVSpeechSynthesisMarkerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWordRange:atByteSampleOffset:"), range_, byteSampleOffset)
	return AVSpeechSynthesisMarkerFromID(rv)
}

// Creates a marker with a type and location of the request’s text.
//
// type: The type that describes the text.
//
// range: The location and length of the request’s text.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(markerType:forTextRange:atByteSampleOffset:)
func (s AVSpeechSynthesisMarker) InitWithMarkerTypeForTextRangeAtByteSampleOffset(type_ AVSpeechSynthesisMarkerMark, range_ foundation.NSRange, byteSampleOffset uint) AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](s.ID, objc.Sel("initWithMarkerType:forTextRange:atByteSampleOffset:"), type_, range_, byteSampleOffset)
	return rv
}
// Creates a word marker with a range of the word and offset into the audio
// buffer.
//
// range: The location and length of the word.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(wordRange:atByteSampleOffset:)
func (s AVSpeechSynthesisMarker) InitWithWordRangeAtByteSampleOffset(range_ foundation.NSRange, byteSampleOffset int) AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](s.ID, objc.Sel("initWithWordRange:atByteSampleOffset:"), range_, byteSampleOffset)
	return rv
}
// Creates a sentence marker with a range of the sentence and offset into the
// audio buffer.
//
// range: The location and length of the word.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(sentenceRange:atByteSampleOffset:)
func (s AVSpeechSynthesisMarker) InitWithSentenceRangeAtByteSampleOffset(range_ foundation.NSRange, byteSampleOffset int) AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](s.ID, objc.Sel("initWithSentenceRange:atByteSampleOffset:"), range_, byteSampleOffset)
	return rv
}
// Creates a paragraph marker with a range of the paragraph and offset into
// the audio buffer.
//
// range: The location and length of the paragraph.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(paragraphRange:atByteSampleOffset:)
func (s AVSpeechSynthesisMarker) InitWithParagraphRangeAtByteSampleOffset(range_ foundation.NSRange, byteSampleOffset int) AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](s.ID, objc.Sel("initWithParagraphRange:atByteSampleOffset:"), range_, byteSampleOffset)
	return rv
}
// Creates a phoneme marker with a range of the phoneme and offset into the
// audio buffer.
//
// phoneme: A string that represents a distinct sound.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(phonemeString:atByteSampleOffset:)
func (s AVSpeechSynthesisMarker) InitWithPhonemeStringAtByteSampleOffset(phoneme string, byteSampleOffset int) AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](s.ID, objc.Sel("initWithPhonemeString:atByteSampleOffset:"), objc.String(phoneme), byteSampleOffset)
	return rv
}
// Creates a bookmark marker with a name and offset into the audio buffer.
//
// mark: The name of the bookmark.
//
// byteSampleOffset: The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/init(bookmarkName:atByteSampleOffset:)
func (s AVSpeechSynthesisMarker) InitWithBookmarkNameAtByteSampleOffset(mark string, byteSampleOffset int) AVSpeechSynthesisMarker {
	rv := objc.Send[AVSpeechSynthesisMarker](s.ID, objc.Sel("initWithBookmarkName:atByteSampleOffset:"), objc.String(mark), byteSampleOffset)
	return rv
}
func (s AVSpeechSynthesisMarker) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The type that describes the text.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/mark-swift.property
func (s AVSpeechSynthesisMarker) Mark() AVSpeechSynthesisMarkerMark {
	rv := objc.Send[AVSpeechSynthesisMarkerMark](s.ID, objc.Sel("mark"))
	return AVSpeechSynthesisMarkerMark(rv)
}
func (s AVSpeechSynthesisMarker) SetMark(value AVSpeechSynthesisMarkerMark) {
	objc.Send[struct{}](s.ID, objc.Sel("setMark:"), value)
}
// A string that represents the name of a bookmark.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/bookmarkName
func (s AVSpeechSynthesisMarker) BookmarkName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("bookmarkName"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisMarker) SetBookmarkName(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setBookmarkName:"), objc.String(value))
}
// A string that represents a distinct sound.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/phoneme
func (s AVSpeechSynthesisMarker) Phoneme() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("phoneme"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesisMarker) SetPhoneme(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setPhoneme:"), objc.String(value))
}
// The location and length of the request’s text.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/textRange
func (s AVSpeechSynthesisMarker) TextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](s.ID, objc.Sel("textRange"))
	return foundation.NSRange(rv)
}
func (s AVSpeechSynthesisMarker) SetTextRange(value foundation.NSRange) {
	objc.Send[struct{}](s.ID, objc.Sel("setTextRange:"), value)
}
// The byte offset into the audio buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesisMarker/byteSampleOffset
func (s AVSpeechSynthesisMarker) ByteSampleOffset() uint {
	rv := objc.Send[uint](s.ID, objc.Sel("byteSampleOffset"))
	return rv
}
func (s AVSpeechSynthesisMarker) SetByteSampleOffset(value uint) {
	objc.Send[struct{}](s.ID, objc.Sel("setByteSampleOffset:"), value)
}
// A block that subclasses use to send marker information to the host.
//
// See: https://developer.apple.com/documentation/avfaudio/avspeechsynthesisprovideraudiounit/speechsynthesisoutputmetadatablock
func (s AVSpeechSynthesisMarker) SpeechSynthesisOutputMetadataBlock() AVSpeechSynthesisProviderOutputBlock {
	rv := objc.Send[AVSpeechSynthesisProviderOutputBlock](s.ID, objc.Sel("speechSynthesisOutputMetadataBlock"))
	return AVSpeechSynthesisProviderOutputBlock(rv)
}
func (s AVSpeechSynthesisMarker) SetSpeechSynthesisOutputMetadataBlock(value AVSpeechSynthesisProviderOutputBlock) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpeechSynthesisOutputMetadataBlock:"), value)
}

