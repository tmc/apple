// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioApplication] class.
var (
	_AVAudioApplicationClass     AVAudioApplicationClass
	_AVAudioApplicationClassOnce sync.Once
)

func getAVAudioApplicationClass() AVAudioApplicationClass {
	_AVAudioApplicationClassOnce.Do(func() {
		_AVAudioApplicationClass = AVAudioApplicationClass{class: objc.GetClass("AVAudioApplication")}
	})
	return _AVAudioApplicationClass
}

// GetAVAudioApplicationClass returns the class object for AVAudioApplication.
func GetAVAudioApplicationClass() AVAudioApplicationClass {
	return getAVAudioApplicationClass()
}

type AVAudioApplicationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioApplicationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioApplicationClass) Alloc() AVAudioApplication {
	rv := objc.Send[AVAudioApplication](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioApplication.ClientID]
//   - [AVAudioApplication.PostNotificationNameUserInfo]
//   - [AVAudioApplication.PrivateCallInputMuteHandlerBlockInputMutedIsTopDownMuteContext]
//   - [AVAudioApplication.PrivateCreateAudioApplicationInServer]
//   - [AVAudioApplication.PrivateEnableSystemMute]
//   - [AVAudioApplication.PrivateHandlePing]
//   - [AVAudioApplication.PrivateOptInToStemClickMuting]
//   - [AVAudioApplication.PrivateRecreateAudioApplicationInServer]
//   - [AVAudioApplication.PrivateSetAppPropertyValue]
//   - [AVAudioApplication.PrivateSetAppPropertyValueGuard]
//   - [AVAudioApplication.PrivateSetInputMuteStateChangeHandlerError]
//   - [AVAudioApplication.PrivateSetInputMutedContextError]
//   - [AVAudioApplication.PrivateSetInputMutedPrimaryOrDelegateContextError]
//   - [AVAudioApplication.PrivateSetInputMutedProxyError]
//   - [AVAudioApplication.PrivateSetMXPropertyOnAllSessionsValue]
//   - [AVAudioApplication.PrivateUpdateAppPropertyValueContext]
//   - [AVAudioApplication.RequestRecordPermissionWithCompletionHandler]
//   - [AVAudioApplication.SessionIDs]
//   - [AVAudioApplication.SetInputMutedContextError]
//   - [AVAudioApplication.StemClickMutingEnabled]
//   - [AVAudioApplication.InitPrivate]
//   - [AVAudioApplication.InitWithSpecification]
//   - [AVAudioApplication.InputMuted]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication
type AVAudioApplication struct {
	objectivec.Object
}

// AVAudioApplicationFromID constructs a [AVAudioApplication] from an objc.ID.
func AVAudioApplicationFromID(id objc.ID) AVAudioApplication {
	return AVAudioApplication{objectivec.Object{ID: id}}
}
// Ensure AVAudioApplication implements IAVAudioApplication.
var _ IAVAudioApplication = AVAudioApplication{}

// An interface definition for the [AVAudioApplication] class.
//
// # Methods
//
//   - [IAVAudioApplication.ClientID]
//   - [IAVAudioApplication.PostNotificationNameUserInfo]
//   - [IAVAudioApplication.PrivateCallInputMuteHandlerBlockInputMutedIsTopDownMuteContext]
//   - [IAVAudioApplication.PrivateCreateAudioApplicationInServer]
//   - [IAVAudioApplication.PrivateEnableSystemMute]
//   - [IAVAudioApplication.PrivateHandlePing]
//   - [IAVAudioApplication.PrivateOptInToStemClickMuting]
//   - [IAVAudioApplication.PrivateRecreateAudioApplicationInServer]
//   - [IAVAudioApplication.PrivateSetAppPropertyValue]
//   - [IAVAudioApplication.PrivateSetAppPropertyValueGuard]
//   - [IAVAudioApplication.PrivateSetInputMuteStateChangeHandlerError]
//   - [IAVAudioApplication.PrivateSetInputMutedContextError]
//   - [IAVAudioApplication.PrivateSetInputMutedPrimaryOrDelegateContextError]
//   - [IAVAudioApplication.PrivateSetInputMutedProxyError]
//   - [IAVAudioApplication.PrivateSetMXPropertyOnAllSessionsValue]
//   - [IAVAudioApplication.PrivateUpdateAppPropertyValueContext]
//   - [IAVAudioApplication.RequestRecordPermissionWithCompletionHandler]
//   - [IAVAudioApplication.SessionIDs]
//   - [IAVAudioApplication.SetInputMutedContextError]
//   - [IAVAudioApplication.StemClickMutingEnabled]
//   - [IAVAudioApplication.InitPrivate]
//   - [IAVAudioApplication.InitWithSpecification]
//   - [IAVAudioApplication.InputMuted]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication
type IAVAudioApplication interface {
	objectivec.IObject

	// Topic: Methods

	ClientID() uint32
	PostNotificationNameUserInfo(name objectivec.IObject, info objectivec.IObject)
	PrivateCallInputMuteHandlerBlockInputMutedIsTopDownMuteContext(block VoidHandler, muted bool, mute bool, context objectivec.IObject) objectivec.IObject
	PrivateCreateAudioApplicationInServer(server objectivec.IObject) bool
	PrivateEnableSystemMute(mute bool)
	PrivateHandlePing()
	PrivateOptInToStemClickMuting()
	PrivateRecreateAudioApplicationInServer() bool
	PrivateSetAppPropertyValue(property objectivec.IObject, value objectivec.IObject) int
	PrivateSetAppPropertyValueGuard(property objectivec.IObject, value objectivec.IObject, guard unsafe.Pointer) int
	PrivateSetInputMuteStateChangeHandlerError(handler func()) (bool, error)
	PrivateSetInputMutedContextError(muted bool, context objectivec.IObject) (bool, error)
	PrivateSetInputMutedPrimaryOrDelegateContextError(delegate bool, context objectivec.IObject) (bool, error)
	PrivateSetInputMutedProxyError(proxy bool) (bool, error)
	PrivateSetMXPropertyOnAllSessionsValue(sessions objectivec.IObject, value objectivec.IObject) int
	PrivateUpdateAppPropertyValueContext(property objectivec.IObject, value objectivec.IObject, context objectivec.IObject) int
	RequestRecordPermissionWithCompletionHandler(handler ErrorHandler)
	SessionIDs() objectivec.IObject
	SetInputMutedContextError(muted bool, context objectivec.IObject) (bool, error)
	StemClickMutingEnabled() bool
	InitPrivate(private objectivec.IObject) AVAudioApplication
	InitWithSpecification(specification objectivec.IObject) AVAudioApplication
	InputMuted() bool
}

// Init initializes the instance.
func (a AVAudioApplication) Init() AVAudioApplication {
	rv := objc.Send[AVAudioApplication](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioApplication) Autorelease() AVAudioApplication {
	rv := objc.Send[AVAudioApplication](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioApplication creates a new AVAudioApplication instance.
func NewAVAudioApplication() AVAudioApplication {
	class := getAVAudioApplicationClass()
	rv := objc.Send[AVAudioApplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/initPrivate:
func NewAudioApplicationPrivate(private objectivec.IObject) AVAudioApplication {
	instance := getAVAudioApplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initPrivate:"), private)
	return AVAudioApplicationFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/initWithSpecification:
func NewAudioApplicationWithSpecification(specification objectivec.IObject) AVAudioApplication {
	instance := getAVAudioApplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpecification:"), specification)
	return AVAudioApplicationFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/postNotificationName:userInfo:
func (a AVAudioApplication) PostNotificationNameUserInfo(name objectivec.IObject, info objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("postNotificationName:userInfo:"), name, info)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateCallInputMuteHandlerBlock:inputMuted:isTopDownMute:context:
func (a AVAudioApplication) PrivateCallInputMuteHandlerBlockInputMutedIsTopDownMuteContext(block VoidHandler, muted bool, mute bool, context objectivec.IObject) objectivec.IObject {
_block0, _ := NewVoidBlock(block)
	rv := objc.Send[objc.ID](a.ID, objc.Sel("privateCallInputMuteHandlerBlock:inputMuted:isTopDownMute:context:"), _block0, muted, mute, context)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateCreateAudioApplicationInServer:
func (a AVAudioApplication) PrivateCreateAudioApplicationInServer(server objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("privateCreateAudioApplicationInServer:"), server)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateEnableSystemMute:
func (a AVAudioApplication) PrivateEnableSystemMute(mute bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("privateEnableSystemMute:"), mute)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateHandlePing
func (a AVAudioApplication) PrivateHandlePing() {
	objc.Send[objc.ID](a.ID, objc.Sel("privateHandlePing"))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateOptInToStemClickMuting
func (a AVAudioApplication) PrivateOptInToStemClickMuting() {
	objc.Send[objc.ID](a.ID, objc.Sel("privateOptInToStemClickMuting"))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateRecreateAudioApplicationInServer
func (a AVAudioApplication) PrivateRecreateAudioApplicationInServer() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("privateRecreateAudioApplicationInServer"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateSetAppProperty:value:
func (a AVAudioApplication) PrivateSetAppPropertyValue(property objectivec.IObject, value objectivec.IObject) int {
	rv := objc.Send[int](a.ID, objc.Sel("privateSetAppProperty:value:"), property, value)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateSetAppProperty:value:guard:
func (a AVAudioApplication) PrivateSetAppPropertyValueGuard(property objectivec.IObject, value objectivec.IObject, guard unsafe.Pointer) int {
	rv := objc.Send[int](a.ID, objc.Sel("privateSetAppProperty:value:guard:"), property, value, guard)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateSetInputMuteStateChangeHandler:error:
func (a AVAudioApplication) PrivateSetInputMuteStateChangeHandlerError(handler func()) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("privateSetInputMuteStateChangeHandler:error:"), handler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("privateSetInputMuteStateChangeHandler:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateSetInputMuted:context:error:
func (a AVAudioApplication) PrivateSetInputMutedContextError(muted bool, context objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("privateSetInputMuted:context:error:"), muted, context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("privateSetInputMuted:context:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateSetInputMutedPrimaryOrDelegate:context:error:
func (a AVAudioApplication) PrivateSetInputMutedPrimaryOrDelegateContextError(delegate bool, context objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("privateSetInputMutedPrimaryOrDelegate:context:error:"), delegate, context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("privateSetInputMutedPrimaryOrDelegate:context:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateSetInputMutedProxy:error:
func (a AVAudioApplication) PrivateSetInputMutedProxyError(proxy bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("privateSetInputMutedProxy:error:"), proxy, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("privateSetInputMutedProxy:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateSetMXPropertyOnAllSessions:value:
func (a AVAudioApplication) PrivateSetMXPropertyOnAllSessionsValue(sessions objectivec.IObject, value objectivec.IObject) int {
	rv := objc.Send[int](a.ID, objc.Sel("privateSetMXPropertyOnAllSessions:value:"), sessions, value)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/privateUpdateAppProperty:value:context:
func (a AVAudioApplication) PrivateUpdateAppPropertyValueContext(property objectivec.IObject, value objectivec.IObject, context objectivec.IObject) int {
	rv := objc.Send[int](a.ID, objc.Sel("privateUpdateAppProperty:value:context:"), property, value, context)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/requestRecordPermissionWithCompletionHandler:
func (a AVAudioApplication) RequestRecordPermissionWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("requestRecordPermissionWithCompletionHandler:"), _block0)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/sessionIDs
func (a AVAudioApplication) SessionIDs() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sessionIDs"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/setInputMuted:context:error:
func (a AVAudioApplication) SetInputMutedContextError(muted bool, context objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("setInputMuted:context:error:"), muted, context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setInputMuted:context:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/stemClickMutingEnabled
func (a AVAudioApplication) StemClickMutingEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("stemClickMutingEnabled"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/initPrivate:
func (a AVAudioApplication) InitPrivate(private objectivec.IObject) AVAudioApplication {
	rv := objc.Send[AVAudioApplication](a.ID, objc.Sel("initPrivate:"), private)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/initWithSpecification:
func (a AVAudioApplication) InitWithSpecification(specification objectivec.IObject) AVAudioApplication {
	rv := objc.Send[AVAudioApplication](a.ID, objc.Sel("initWithSpecification:"), specification)
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/allowAppToInitiatePlaybackTemporarily:error:
func (_AVAudioApplicationClass AVAudioApplicationClass) AllowAppToInitiatePlaybackTemporarilyError(temporarily objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("allowAppToInitiatePlaybackTemporarily:error:"), temporarily, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("allowAppToInitiatePlaybackTemporarily:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/appleTVSupportsEnhanceDialogue
func (_AVAudioApplicationClass AVAudioApplicationClass) AppleTVSupportsEnhanceDialogue() bool {
	rv := objc.Send[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("appleTVSupportsEnhanceDialogue"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/currentRouteSupportsEnhanceDialogue:
func (_AVAudioApplicationClass AVAudioApplicationClass) CurrentRouteSupportsEnhanceDialogue(dialogue []objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("currentRouteSupportsEnhanceDialogue:"), objectivec.IObjectSliceToNSArray(dialogue))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/getEnhanceDialogueLevel:error:
func (_AVAudioApplicationClass AVAudioApplicationClass) GetEnhanceDialogueLevelError() (int64, error) {
	var level int64
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("getEnhanceDialogueLevel:error:"), unsafe.Pointer(&level), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0, errors.New("getEnhanceDialogueLevel:error: returned NO with nil NSError")
	}
	return level, nil
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/iosDeviceSupportsEnhanceDialogue
func (_AVAudioApplicationClass AVAudioApplicationClass) IosDeviceSupportsEnhanceDialogue() bool {
	rv := objc.Send[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("iosDeviceSupportsEnhanceDialogue"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/muteRunningInputs:
func (_AVAudioApplicationClass AVAudioApplicationClass) MuteRunningInputs(inputs []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioApplicationClass.class), objc.Sel("muteRunningInputs:"), objectivec.IObjectSliceToNSArray(inputs))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/setEnhanceDialogueLevel:error:
func (_AVAudioApplicationClass AVAudioApplicationClass) SetEnhanceDialogueLevelError(level int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("setEnhanceDialogueLevel:error:"), level, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setEnhanceDialogueLevel:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/setEnhanceDialoguePreference:error:
func (_AVAudioApplicationClass AVAudioApplicationClass) SetEnhanceDialoguePreferenceError(preference int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("setEnhanceDialoguePreference:error:"), preference, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setEnhanceDialoguePreference:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/toggleInputMute:
func (_AVAudioApplicationClass AVAudioApplicationClass) ToggleInputMute(mute []objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("toggleInputMute:"), objectivec.IObjectSliceToNSArray(mute))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/visionosDeviceSupportsEnhanceDialogue
func (_AVAudioApplicationClass AVAudioApplicationClass) VisionosDeviceSupportsEnhanceDialogue() bool {
	rv := objc.Send[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("visionosDeviceSupportsEnhanceDialogue"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/clientID
func (a AVAudioApplication) ClientID() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("clientID"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/inputMuted
func (a AVAudioApplication) InputMuted() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("inputMuted"))
	return rv
}

// RequestRecordPermission is a synchronous wrapper around [AVAudioApplication.RequestRecordPermissionWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioApplication) RequestRecordPermission(ctx context.Context) error {
	done := make(chan error, 1)
	a.RequestRecordPermissionWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

