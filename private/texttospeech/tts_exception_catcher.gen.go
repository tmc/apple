// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSExceptionCatcher] class.
var (
	_TTSExceptionCatcherClass     TTSExceptionCatcherClass
	_TTSExceptionCatcherClassOnce sync.Once
)

func getTTSExceptionCatcherClass() TTSExceptionCatcherClass {
	_TTSExceptionCatcherClassOnce.Do(func() {
		_TTSExceptionCatcherClass = TTSExceptionCatcherClass{class: objc.GetClass("TTSExceptionCatcher")}
	})
	return _TTSExceptionCatcherClass
}

// GetTTSExceptionCatcherClass returns the class object for TTSExceptionCatcher.
func GetTTSExceptionCatcherClass() TTSExceptionCatcherClass {
	return getTTSExceptionCatcherClass()
}

type TTSExceptionCatcherClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSExceptionCatcherClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSExceptionCatcherClass) Alloc() TTSExceptionCatcher {
	rv := objc.Send[TTSExceptionCatcher](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSExceptionCatcher
type TTSExceptionCatcher struct {
	objectivec.Object
}

// TTSExceptionCatcherFromID constructs a [TTSExceptionCatcher] from an objc.ID.
func TTSExceptionCatcherFromID(id objc.ID) TTSExceptionCatcher {
	return TTSExceptionCatcher{objectivec.Object{ID: id}}
}

// Ensure TTSExceptionCatcher implements ITTSExceptionCatcher.
var _ ITTSExceptionCatcher = TTSExceptionCatcher{}

// An interface definition for the [TTSExceptionCatcher] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSExceptionCatcher
type ITTSExceptionCatcher interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSExceptionCatcher) Init() TTSExceptionCatcher {
	rv := objc.Send[TTSExceptionCatcher](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSExceptionCatcher) Autorelease() TTSExceptionCatcher {
	rv := objc.Send[TTSExceptionCatcher](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSExceptionCatcher creates a new TTSExceptionCatcher instance.
func NewTTSExceptionCatcher() TTSExceptionCatcher {
	class := getTTSExceptionCatcherClass()
	rv := objc.Send[TTSExceptionCatcher](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSExceptionCatcher/catchException:error:
func (_TTSExceptionCatcherClass TTSExceptionCatcherClass) CatchExceptionError(exception func()) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_TTSExceptionCatcherClass.class), objc.Sel("catchException:error:"), exception, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("catchException:error: returned NO with nil NSError")
	}
	return rv, nil

}
