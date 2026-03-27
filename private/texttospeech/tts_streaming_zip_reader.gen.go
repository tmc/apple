// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSStreamingZipReader] class.
var (
	_TTSStreamingZipReaderClass     TTSStreamingZipReaderClass
	_TTSStreamingZipReaderClassOnce sync.Once
)

func getTTSStreamingZipReaderClass() TTSStreamingZipReaderClass {
	_TTSStreamingZipReaderClassOnce.Do(func() {
		_TTSStreamingZipReaderClass = TTSStreamingZipReaderClass{class: objc.GetClass("TTSStreamingZipReader")}
	})
	return _TTSStreamingZipReaderClass
}

// GetTTSStreamingZipReaderClass returns the class object for TTSStreamingZipReader.
func GetTTSStreamingZipReaderClass() TTSStreamingZipReaderClass {
	return getTTSStreamingZipReaderClass()
}

type TTSStreamingZipReaderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSStreamingZipReaderClass) Alloc() TTSStreamingZipReader {
	rv := objc.Send[TTSStreamingZipReader](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSStreamingZipReader.EnumerateFiles]
//   - [TTSStreamingZipReader.Password]
//   - [TTSStreamingZipReader.SetPassword]
//   - [TTSStreamingZipReader.ZipPath]
//   - [TTSStreamingZipReader.SetZipPath]
//   - [TTSStreamingZipReader.InitWithPathAndPassword]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStreamingZipReader
type TTSStreamingZipReader struct {
	objectivec.Object
}

// TTSStreamingZipReaderFromID constructs a [TTSStreamingZipReader] from an objc.ID.
func TTSStreamingZipReaderFromID(id objc.ID) TTSStreamingZipReader {
	return TTSStreamingZipReader{objectivec.Object{ID: id}}
}
// Ensure TTSStreamingZipReader implements ITTSStreamingZipReader.
var _ ITTSStreamingZipReader = TTSStreamingZipReader{}

// An interface definition for the [TTSStreamingZipReader] class.
//
// # Methods
//
//   - [ITTSStreamingZipReader.EnumerateFiles]
//   - [ITTSStreamingZipReader.Password]
//   - [ITTSStreamingZipReader.SetPassword]
//   - [ITTSStreamingZipReader.ZipPath]
//   - [ITTSStreamingZipReader.SetZipPath]
//   - [ITTSStreamingZipReader.InitWithPathAndPassword]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStreamingZipReader
type ITTSStreamingZipReader interface {
	objectivec.IObject

	// Topic: Methods

	EnumerateFiles(files VoidHandler) bool
	Password() string
	SetPassword(value string)
	ZipPath() string
	SetZipPath(value string)
	InitWithPathAndPassword(path objectivec.IObject, password objectivec.IObject) TTSStreamingZipReader
}

// Init initializes the instance.
func (t TTSStreamingZipReader) Init() TTSStreamingZipReader {
	rv := objc.Send[TTSStreamingZipReader](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSStreamingZipReader) Autorelease() TTSStreamingZipReader {
	rv := objc.Send[TTSStreamingZipReader](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSStreamingZipReader creates a new TTSStreamingZipReader instance.
func NewTTSStreamingZipReader() TTSStreamingZipReader {
	class := getTTSStreamingZipReaderClass()
	rv := objc.Send[TTSStreamingZipReader](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStreamingZipReader/initWithPath:andPassword:
func NewTTSStreamingZipReaderWithPathAndPassword(path objectivec.IObject, password objectivec.IObject) TTSStreamingZipReader {
	instance := getTTSStreamingZipReaderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPath:andPassword:"), path, password)
	return TTSStreamingZipReaderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStreamingZipReader/enumerateFiles:
func (t TTSStreamingZipReader) EnumerateFiles(files VoidHandler) bool {
_block0, _cleanup0 := NewVoidBlock(files)
	defer _cleanup0()
	rv := objc.Send[bool](t.ID, objc.Sel("enumerateFiles:"), _block0)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStreamingZipReader/initWithPath:andPassword:
func (t TTSStreamingZipReader) InitWithPathAndPassword(path objectivec.IObject, password objectivec.IObject) TTSStreamingZipReader {
	rv := objc.Send[TTSStreamingZipReader](t.ID, objc.Sel("initWithPath:andPassword:"), path, password)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSStreamingZipReader/password
func (t TTSStreamingZipReader) Password() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("password"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSStreamingZipReader) SetPassword(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPassword:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSStreamingZipReader/zipPath
func (t TTSStreamingZipReader) ZipPath() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("zipPath"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSStreamingZipReader) SetZipPath(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setZipPath:"), objc.String(value))
}

// EnumerateFilesSync is a synchronous wrapper around [TTSStreamingZipReader.EnumerateFiles].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSStreamingZipReader) EnumerateFilesSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.EnumerateFiles(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

