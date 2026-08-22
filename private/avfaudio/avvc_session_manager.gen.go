// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCSessionManager] class.
var (
	_AVVCSessionManagerClass     AVVCSessionManagerClass
	_AVVCSessionManagerClassOnce sync.Once
)

func getAVVCSessionManagerClass() AVVCSessionManagerClass {
	_AVVCSessionManagerClassOnce.Do(func() {
		_AVVCSessionManagerClass = AVVCSessionManagerClass{class: objc.GetClass("AVVCSessionManager")}
	})
	return _AVVCSessionManagerClass
}

// GetAVVCSessionManagerClass returns the class object for AVVCSessionManager.
func GetAVVCSessionManagerClass() AVVCSessionManagerClass {
	return getAVVCSessionManagerClass()
}

type AVVCSessionManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCSessionManagerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCSessionManagerClass) Alloc() AVVCSessionManager {
	rv := objc.SendIfResponds[AVVCSessionManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCSessionManager.AudioSession]
//   - [AVVCSessionManager.SetAudioSession]
//   - [AVVCSessionManager.IsCurrentInputBuiltInMic]
//   - [AVVCSessionManager.SetIsUsingBuiltInMicForRecordingError]
//   - [AVVCSessionManager.SetupOneTimeSessionSettingsForClient]
//   - [AVVCSessionManager.InitWithSession]
type AVVCSessionManager struct {
	objectivec.Object
}

// AVVCSessionManagerFromID constructs a [AVVCSessionManager] from an objc.ID.
func AVVCSessionManagerFromID(id objc.ID) AVVCSessionManager {
	return AVVCSessionManager{objectivec.Object{ID: id}}
}

// Ensure AVVCSessionManager implements IAVVCSessionManager.
var _ IAVVCSessionManager = AVVCSessionManager{}

// An interface definition for the [AVVCSessionManager] class.
//
// # Methods
//
//   - [IAVVCSessionManager.AudioSession]
//   - [IAVVCSessionManager.SetAudioSession]
//   - [IAVVCSessionManager.IsCurrentInputBuiltInMic]
//   - [IAVVCSessionManager.SetIsUsingBuiltInMicForRecordingError]
//   - [IAVVCSessionManager.SetupOneTimeSessionSettingsForClient]
//   - [IAVVCSessionManager.InitWithSession]
type IAVVCSessionManager interface {
	objectivec.IObject

	// Topic: Methods

	AudioSession() objectivec.IObject
	SetAudioSession(value objectivec.IObject)
	IsCurrentInputBuiltInMic() bool
	SetIsUsingBuiltInMicForRecordingError(recording bool) (bool, error)
	SetupOneTimeSessionSettingsForClient(client int64) int
	InitWithSession(session objectivec.IObject) AVVCSessionManager
}

// Init initializes the instance.
func (a AVVCSessionManager) Init() AVVCSessionManager {
	rv := objc.SendIfResponds[AVVCSessionManager](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCSessionManager) Autorelease() AVVCSessionManager {
	rv := objc.SendIfResponds[AVVCSessionManager](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCSessionManager creates a new AVVCSessionManager instance.
func NewAVVCSessionManager() AVVCSessionManager {
	class := getAVVCSessionManagerClass()
	rv := objc.SendIfResponds[AVVCSessionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVCSessionManagerWithSession(session objectivec.IObject) AVVCSessionManager {
	instance := getAVVCSessionManagerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return AVVCSessionManagerFromID(rv)
}

func (a AVVCSessionManager) IsCurrentInputBuiltInMic() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("isCurrentInputBuiltInMic"))
	return rv
}
func (a AVVCSessionManager) SetIsUsingBuiltInMicForRecordingError(recording bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("setIsUsingBuiltInMicForRecording:error:"), recording, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setIsUsingBuiltInMicForRecording:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a AVVCSessionManager) SetupOneTimeSessionSettingsForClient(client int64) int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("setupOneTimeSessionSettingsForClient:"), client)
	return rv
}
func (a AVVCSessionManager) InitWithSession(session objectivec.IObject) AVVCSessionManager {
	rv := objc.SendIfResponds[AVVCSessionManager](a.ID, objc.Sel("initWithSession:"), session)
	return rv
}

func (a AVVCSessionManager) AudioSession() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("audioSession"))
	return objectivec.Object{ID: rv}
}
func (a AVVCSessionManager) SetAudioSession(value objectivec.IObject) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setAudioSession:"), value)
}
