// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UtteranceRunner] class.
var (
	_UtteranceRunnerClass     UtteranceRunnerClass
	_UtteranceRunnerClassOnce sync.Once
)

func getUtteranceRunnerClass() UtteranceRunnerClass {
	_UtteranceRunnerClassOnce.Do(func() {
		_UtteranceRunnerClass = UtteranceRunnerClass{class: objc.GetClass("_TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner")}
	})
	return _UtteranceRunnerClass
}

// GetUtteranceRunnerClass returns the class object for _TtCC12TextToSpeech15CoreSynthesizer15UtteranceRunner.
func GetUtteranceRunnerClass() UtteranceRunnerClass {
	return getUtteranceRunnerClass()
}

type UtteranceRunnerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UtteranceRunnerClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UtteranceRunnerClass) Alloc() UtteranceRunner {
	rv := objc.SendIfResponds[UtteranceRunner](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

type UtteranceRunner struct {
	objectivec.Object
}

// UtteranceRunnerFromID constructs a [UtteranceRunner] from an objc.ID.
func UtteranceRunnerFromID(id objc.ID) UtteranceRunner {
	return UtteranceRunner{objectivec.Object{ID: id}}
}

// Ensure UtteranceRunner implements IUtteranceRunner.
var _ IUtteranceRunner = UtteranceRunner{}

// An interface definition for the [UtteranceRunner] class.
type IUtteranceRunner interface {
	objectivec.IObject
}

// Init initializes the instance.
func (u UtteranceRunner) Init() UtteranceRunner {
	rv := objc.SendIfResponds[UtteranceRunner](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UtteranceRunner) Autorelease() UtteranceRunner {
	rv := objc.SendIfResponds[UtteranceRunner](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUtteranceRunner creates a new UtteranceRunner instance.
func NewUtteranceRunner() UtteranceRunner {
	class := getUtteranceRunnerClass()
	rv := objc.SendIfResponds[UtteranceRunner](objc.ID(class.class), objc.Sel("new"))
	return rv
}
