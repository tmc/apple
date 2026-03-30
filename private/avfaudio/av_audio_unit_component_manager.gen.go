// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitComponentManager] class.
var (
	_AVAudioUnitComponentManagerClass     AVAudioUnitComponentManagerClass
	_AVAudioUnitComponentManagerClassOnce sync.Once
)

func getAVAudioUnitComponentManagerClass() AVAudioUnitComponentManagerClass {
	_AVAudioUnitComponentManagerClassOnce.Do(func() {
		_AVAudioUnitComponentManagerClass = AVAudioUnitComponentManagerClass{class: objc.GetClass("AVAudioUnitComponentManager")}
	})
	return _AVAudioUnitComponentManagerClass
}

// GetAVAudioUnitComponentManagerClass returns the class object for AVAudioUnitComponentManager.
func GetAVAudioUnitComponentManagerClass() AVAudioUnitComponentManagerClass {
	return getAVAudioUnitComponentManagerClass()
}

type AVAudioUnitComponentManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitComponentManagerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitComponentManagerClass) Alloc() AVAudioUnitComponentManager {
	rv := objc.Send[AVAudioUnitComponentManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioUnitComponentManager.LocaleChanged]
//   - [AVAudioUnitComponentManager.RegistrationsChanged]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager
type AVAudioUnitComponentManager struct {
	objectivec.Object
}

// AVAudioUnitComponentManagerFromID constructs a [AVAudioUnitComponentManager] from an objc.ID.
func AVAudioUnitComponentManagerFromID(id objc.ID) AVAudioUnitComponentManager {
	return AVAudioUnitComponentManager{objectivec.Object{ID: id}}
}

// Ensure AVAudioUnitComponentManager implements IAVAudioUnitComponentManager.
var _ IAVAudioUnitComponentManager = AVAudioUnitComponentManager{}

// An interface definition for the [AVAudioUnitComponentManager] class.
//
// # Methods
//
//   - [IAVAudioUnitComponentManager.LocaleChanged]
//   - [IAVAudioUnitComponentManager.RegistrationsChanged]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager
type IAVAudioUnitComponentManager interface {
	objectivec.IObject

	// Topic: Methods

	LocaleChanged(changed objectivec.IObject)
	RegistrationsChanged(changed objectivec.IObject)
}

// Init initializes the instance.
func (a AVAudioUnitComponentManager) Init() AVAudioUnitComponentManager {
	rv := objc.Send[AVAudioUnitComponentManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitComponentManager) Autorelease() AVAudioUnitComponentManager {
	rv := objc.Send[AVAudioUnitComponentManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitComponentManager creates a new AVAudioUnitComponentManager instance.
func NewAVAudioUnitComponentManager() AVAudioUnitComponentManager {
	class := getAVAudioUnitComponentManagerClass()
	rv := objc.Send[AVAudioUnitComponentManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/localeChanged:
func (a AVAudioUnitComponentManager) LocaleChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("localeChanged:"), changed)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/registrationsChanged:
func (a AVAudioUnitComponentManager) RegistrationsChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("registrationsChanged:"), changed)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/privateAllocInitSingleton
func (_AVAudioUnitComponentManagerClass AVAudioUnitComponentManagerClass) PrivateAllocInitSingleton() {
	objc.Send[objc.ID](objc.ID(_AVAudioUnitComponentManagerClass.class), objc.Sel("privateAllocInitSingleton"))
}
