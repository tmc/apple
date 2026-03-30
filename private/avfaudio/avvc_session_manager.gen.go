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
	rv := objc.Send[AVVCSessionManager](objc.ID(ac.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionManager
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
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionManager
type IAVVCSessionManager interface {
	objectivec.IObject

	// Topic: Methods

	AudioSession() objc.ID
	SetAudioSession(value objc.ID)
	IsCurrentInputBuiltInMic() bool
	SetIsUsingBuiltInMicForRecordingError(recording bool) (bool, error)
	SetupOneTimeSessionSettingsForClient(client int64) int
	InitWithSession(session objectivec.IObject) AVVCSessionManager
}

// Init initializes the instance.
func (v AVVCSessionManager) Init() AVVCSessionManager {
	rv := objc.Send[AVVCSessionManager](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCSessionManager) Autorelease() AVVCSessionManager {
	rv := objc.Send[AVVCSessionManager](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCSessionManager creates a new AVVCSessionManager instance.
func NewAVVCSessionManager() AVVCSessionManager {
	class := getAVVCSessionManagerClass()
	rv := objc.Send[AVVCSessionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionManager/initWithSession:
func NewVCSessionManagerWithSession(session objectivec.IObject) AVVCSessionManager {
	instance := getAVVCSessionManagerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSession:"), session)
	return AVVCSessionManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionManager/isCurrentInputBuiltInMic
func (v AVVCSessionManager) IsCurrentInputBuiltInMic() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isCurrentInputBuiltInMic"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionManager/setIsUsingBuiltInMicForRecording:error:
func (v AVVCSessionManager) SetIsUsingBuiltInMicForRecordingError(recording bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("setIsUsingBuiltInMicForRecording:error:"), recording, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setIsUsingBuiltInMicForRecording:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionManager/setupOneTimeSessionSettingsForClient:
func (v AVVCSessionManager) SetupOneTimeSessionSettingsForClient(client int64) int {
	rv := objc.Send[int](v.ID, objc.Sel("setupOneTimeSessionSettingsForClient:"), client)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionManager/initWithSession:
func (v AVVCSessionManager) InitWithSession(session objectivec.IObject) AVVCSessionManager {
	rv := objc.Send[AVVCSessionManager](v.ID, objc.Sel("initWithSession:"), session)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCSessionManager/audioSession
func (v AVVCSessionManager) AudioSession() objc.ID {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("audioSession"))
	return rv
}
func (v AVVCSessionManager) SetAudioSession(value objc.ID) {
	objc.Send[struct{}](v.ID, objc.Sel("setAudioSession:"), value)
}
