// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSUnicodeUtils] class.
var (
	_TTSUnicodeUtilsClass     TTSUnicodeUtilsClass
	_TTSUnicodeUtilsClassOnce sync.Once
)

func getTTSUnicodeUtilsClass() TTSUnicodeUtilsClass {
	_TTSUnicodeUtilsClassOnce.Do(func() {
		_TTSUnicodeUtilsClass = TTSUnicodeUtilsClass{class: objc.GetClass("TTSUnicodeUtils")}
	})
	return _TTSUnicodeUtilsClass
}

// GetTTSUnicodeUtilsClass returns the class object for TTSUnicodeUtils.
func GetTTSUnicodeUtilsClass() TTSUnicodeUtilsClass {
	return getTTSUnicodeUtilsClass()
}

type TTSUnicodeUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSUnicodeUtilsClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSUnicodeUtilsClass) Alloc() TTSUnicodeUtils {
	rv := objc.Send[TTSUnicodeUtils](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSUnicodeUtils
type TTSUnicodeUtils struct {
	objectivec.Object
}

// TTSUnicodeUtilsFromID constructs a [TTSUnicodeUtils] from an objc.ID.
func TTSUnicodeUtilsFromID(id objc.ID) TTSUnicodeUtils {
	return TTSUnicodeUtils{objectivec.Object{ID: id}}
}
// Ensure TTSUnicodeUtils implements ITTSUnicodeUtils.
var _ ITTSUnicodeUtils = TTSUnicodeUtils{}

// An interface definition for the [TTSUnicodeUtils] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSUnicodeUtils
type ITTSUnicodeUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSUnicodeUtils) Init() TTSUnicodeUtils {
	rv := objc.Send[TTSUnicodeUtils](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSUnicodeUtils) Autorelease() TTSUnicodeUtils {
	rv := objc.Send[TTSUnicodeUtils](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSUnicodeUtils creates a new TTSUnicodeUtils instance.
func NewTTSUnicodeUtils() TTSUnicodeUtils {
	class := getTTSUnicodeUtilsClass()
	rv := objc.Send[TTSUnicodeUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSUnicodeUtils/codePointToUtf8ByteSize:
func (_TTSUnicodeUtilsClass TTSUnicodeUtilsClass) CodePointToUtf8ByteSize(size uint32) byte {
	rv := objc.Send[byte](objc.ID(_TTSUnicodeUtilsClass.class), objc.Sel("codePointToUtf8ByteSize:"), size)
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSUnicodeUtils/utf16RangeFromUTF8Range:chars:size:
func (_TTSUnicodeUtilsClass TTSUnicodeUtilsClass) Utf16RangeFromUTF8RangeCharsSize(uTF8Range foundation.NSRange, chars string, size uint64) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](objc.ID(_TTSUnicodeUtilsClass.class), objc.Sel("utf16RangeFromUTF8Range:chars:size:"), uTF8Range, unsafe.Pointer(unsafe.StringData(chars + "\x00")), size)
	return foundation.NSRange(rv)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSUnicodeUtils/utf8RangeFromUTF16Range:chars:size:
func (_TTSUnicodeUtilsClass TTSUnicodeUtilsClass) Utf8RangeFromUTF16RangeCharsSize(uTF16Range foundation.NSRange, chars unsafe.Pointer, size uint64) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](objc.ID(_TTSUnicodeUtilsClass.class), objc.Sel("utf8RangeFromUTF16Range:chars:size:"), uTF16Range, chars, size)
	return foundation.NSRange(rv)
}

