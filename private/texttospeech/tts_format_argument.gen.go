// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSFormatArgument] class.
var (
	_TTSFormatArgumentClass     TTSFormatArgumentClass
	_TTSFormatArgumentClassOnce sync.Once
)

func getTTSFormatArgumentClass() TTSFormatArgumentClass {
	_TTSFormatArgumentClassOnce.Do(func() {
		_TTSFormatArgumentClass = TTSFormatArgumentClass{class: objc.GetClass("TTSFormatArgument")}
	})
	return _TTSFormatArgumentClass
}

// GetTTSFormatArgumentClass returns the class object for TTSFormatArgument.
func GetTTSFormatArgumentClass() TTSFormatArgumentClass {
	return getTTSFormatArgumentClass()
}

type TTSFormatArgumentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSFormatArgumentClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSFormatArgumentClass) Alloc() TTSFormatArgument {
	rv := objc.Send[TTSFormatArgument](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSFormatArgument
type TTSFormatArgument struct {
	objectivec.Object
}

// TTSFormatArgumentFromID constructs a [TTSFormatArgument] from an objc.ID.
func TTSFormatArgumentFromID(id objc.ID) TTSFormatArgument {
	return TTSFormatArgument{objectivec.Object{ID: id}}
}

// Ensure TTSFormatArgument implements ITTSFormatArgument.
var _ ITTSFormatArgument = TTSFormatArgument{}

// An interface definition for the [TTSFormatArgument] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSFormatArgument
type ITTSFormatArgument interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSFormatArgument) Init() TTSFormatArgument {
	rv := objc.Send[TTSFormatArgument](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSFormatArgument) Autorelease() TTSFormatArgument {
	rv := objc.Send[TTSFormatArgument](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSFormatArgument creates a new TTSFormatArgument instance.
func NewTTSFormatArgument() TTSFormatArgument {
	class := getTTSFormatArgumentClass()
	rv := objc.Send[TTSFormatArgument](objc.ID(class.class), objc.Sel("new"))
	return rv
}
