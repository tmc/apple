// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner] class.
var (
	_TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass     TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass
	_TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClassOnce sync.Once
)

func getTtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass() TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass {
	_TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClassOnce.Do(func() {
		_TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass = TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass{class: objc.GetClass("_TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner")}
	})
	return _TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass
}

// GetTtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass returns the class object for _TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner.
func GetTtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass() TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass {
	return getTtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass()
}

type TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass) Alloc() TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner {
	rv := objc.Send[TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/_TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner
type TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner struct {
	objectivec.Object
}

// TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerFromID constructs a [TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner] from an objc.ID.
func TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerFromID(id objc.ID) TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner {
	return TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner{objectivec.Object{ID: id}}
}
// NOTE: TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner struct embeds objectivec.Object (parent type unavailable) but
// ITtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/_TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner
type ITtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner) Init() TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner {
	rv := objc.Send[TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner) Autorelease() TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner {
	rv := objc.Send[TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner creates a new TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner instance.
func NewTtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner() TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner {
	class := getTtCC12TextToSpeech15CoreSynthesizer15UtteranceRunnerClass()
	rv := objc.Send[TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner](objc.ID(class.class), objc.Sel("new"))
	return rv
}

