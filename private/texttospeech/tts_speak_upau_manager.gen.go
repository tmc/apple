// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSSpeakUPAUManager] class.
var (
	_TTSSpeakUPAUManagerClass     TTSSpeakUPAUManagerClass
	_TTSSpeakUPAUManagerClassOnce sync.Once
)

func getTTSSpeakUPAUManagerClass() TTSSpeakUPAUManagerClass {
	_TTSSpeakUPAUManagerClassOnce.Do(func() {
		_TTSSpeakUPAUManagerClass = TTSSpeakUPAUManagerClass{class: objc.GetClass("TTSSpeakUPAUManager")}
	})
	return _TTSSpeakUPAUManagerClass
}

// GetTTSSpeakUPAUManagerClass returns the class object for TTSSpeakUPAUManager.
func GetTTSSpeakUPAUManagerClass() TTSSpeakUPAUManagerClass {
	return getTTSSpeakUPAUManagerClass()
}

type TTSSpeakUPAUManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSSpeakUPAUManagerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSpeakUPAUManagerClass) Alloc() TTSSpeakUPAUManager {
	rv := objc.Send[TTSSpeakUPAUManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeakUPAUManager
type TTSSpeakUPAUManager struct {
	objectivec.Object
}

// TTSSpeakUPAUManagerFromID constructs a [TTSSpeakUPAUManager] from an objc.ID.
func TTSSpeakUPAUManagerFromID(id objc.ID) TTSSpeakUPAUManager {
	return TTSSpeakUPAUManager{objectivec.Object{ID: id}}
}
// Ensure TTSSpeakUPAUManager implements ITTSSpeakUPAUManager.
var _ ITTSSpeakUPAUManager = TTSSpeakUPAUManager{}

// An interface definition for the [TTSSpeakUPAUManager] class.
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeakUPAUManager
type ITTSSpeakUPAUManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TTSSpeakUPAUManager) Init() TTSSpeakUPAUManager {
	rv := objc.Send[TTSSpeakUPAUManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSpeakUPAUManager) Autorelease() TTSSpeakUPAUManager {
	rv := objc.Send[TTSSpeakUPAUManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSpeakUPAUManager creates a new TTSSpeakUPAUManager instance.
func NewTTSSpeakUPAUManager() TTSSpeakUPAUManager {
	class := getTTSSpeakUPAUManagerClass()
	rv := objc.Send[TTSSpeakUPAUManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeakUPAUManager/component
func (_TTSSpeakUPAUManagerClass TTSSpeakUPAUManagerClass) Component() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeakUPAUManagerClass.class), objc.Sel("component"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeakUPAUManager/registerAU
func (_TTSSpeakUPAUManagerClass TTSSpeakUPAUManagerClass) RegisterAU() {
	objc.Send[objc.ID](objc.ID(_TTSSpeakUPAUManagerClass.class), objc.Sel("registerAU"))
}

