// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[AVAudioApplication](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioApplication.ClientID]
//   - [AVAudioApplication.PostNotificationNameUserInfo]
//   - [AVAudioApplication.PrivateCallInputMuteHandlerBlockInputMutedIsTopDownMuteContext]
//   - [AVAudioApplication.PrivateCreateAudioApplicationInServer]
//   - [AVAudioApplication.PrivateEnableSystemMute]
//   - [AVAudioApplication.PrivateGetAppProperty]
//   - [AVAudioApplication.PrivateGetMXProperty]
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
//   - [AVAudioApplication.InitDelegateForProcessProcessAttribution]
//   - [AVAudioApplication.InitPrivate]
//   - [AVAudioApplication.InitProxyForProcess]
//   - [AVAudioApplication.InitWithSpecification]
//   - [AVAudioApplication.InputMuted]
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
//   - [IAVAudioApplication.PrivateGetAppProperty]
//   - [IAVAudioApplication.PrivateGetMXProperty]
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
//   - [IAVAudioApplication.InitDelegateForProcessProcessAttribution]
//   - [IAVAudioApplication.InitPrivate]
//   - [IAVAudioApplication.InitProxyForProcess]
//   - [IAVAudioApplication.InitWithSpecification]
//   - [IAVAudioApplication.InputMuted]
type IAVAudioApplication interface {
	objectivec.IObject

	// Topic: Methods

	ClientID() uint32
	PostNotificationNameUserInfo(name objectivec.IObject, info objectivec.IObject)
	PrivateCallInputMuteHandlerBlockInputMutedIsTopDownMuteContext(block VoidHandler, muted bool, mute bool, context objectivec.IObject) objectivec.IObject
	PrivateCreateAudioApplicationInServer(server objectivec.IObject) bool
	PrivateEnableSystemMute(mute bool)
	PrivateGetAppProperty(property objectivec.IObject) unsafe.Pointer
	PrivateGetMXProperty(mXProperty objectivec.IObject) unsafe.Pointer
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
	InitDelegateForProcessProcessAttribution(process unsafe.Pointer, attribution objectivec.IObject) AVAudioApplication
	InitPrivate(private objectivec.IObject) AVAudioApplication
	InitProxyForProcess(process unsafe.Pointer) AVAudioApplication
	InitWithSpecification(specification objectivec.IObject) AVAudioApplication
	InputMuted() bool
}

// Init initializes the instance.
func (a AVAudioApplication) Init() AVAudioApplication {
	rv := objc.SendIfResponds[AVAudioApplication](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioApplication) Autorelease() AVAudioApplication {
	rv := objc.SendIfResponds[AVAudioApplication](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioApplication creates a new AVAudioApplication instance.
func NewAVAudioApplication() AVAudioApplication {
	class := getAVAudioApplicationClass()
	rv := objc.SendIfResponds[AVAudioApplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAudioApplicationDelegateForProcessProcessAttribution(process unsafe.Pointer, attribution objectivec.IObject) AVAudioApplication {
	instance := getAVAudioApplicationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initDelegateForProcess:processAttribution:"), process, attribution)
	return AVAudioApplicationFromID(rv)
}

func NewAudioApplicationPrivate(private objectivec.IObject) AVAudioApplication {
	instance := getAVAudioApplicationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initPrivate:"), private)
	return AVAudioApplicationFromID(rv)
}

func NewAudioApplicationProxyForProcess(process unsafe.Pointer) AVAudioApplication {
	instance := getAVAudioApplicationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initProxyForProcess:"), process)
	return AVAudioApplicationFromID(rv)
}

func NewAudioApplicationWithSpecification(specification objectivec.IObject) AVAudioApplication {
	instance := getAVAudioApplicationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpecification:"), specification)
	return AVAudioApplicationFromID(rv)
}

func (a AVAudioApplication) PostNotificationNameUserInfo(name objectivec.IObject, info objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("postNotificationName:userInfo:"), name, info)
}
func (a AVAudioApplication) PrivateCallInputMuteHandlerBlockInputMutedIsTopDownMuteContext(block VoidHandler, muted bool, mute bool, context objectivec.IObject) objectivec.IObject {
	_block0, _ := NewVoidBlock(block)
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("privateCallInputMuteHandlerBlock:inputMuted:isTopDownMute:context:"), _block0, muted, mute, context)
	return objectivec.Object{ID: rv}
}
func (a AVAudioApplication) PrivateCreateAudioApplicationInServer(server objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("privateCreateAudioApplicationInServer:"), server)
	return rv
}
func (a AVAudioApplication) PrivateEnableSystemMute(mute bool) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("privateEnableSystemMute:"), mute)
}
func (a AVAudioApplication) PrivateGetAppProperty(property objectivec.IObject) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](a.ID, objc.Sel("privateGetAppProperty:"), property)
	return rv
}
func (a AVAudioApplication) PrivateGetMXProperty(mXProperty objectivec.IObject) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](a.ID, objc.Sel("privateGetMXProperty:"), mXProperty)
	return rv
}
func (a AVAudioApplication) PrivateHandlePing() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("privateHandlePing"))
}
func (a AVAudioApplication) PrivateOptInToStemClickMuting() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("privateOptInToStemClickMuting"))
}
func (a AVAudioApplication) PrivateRecreateAudioApplicationInServer() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("privateRecreateAudioApplicationInServer"))
	return rv
}
func (a AVAudioApplication) PrivateSetAppPropertyValue(property objectivec.IObject, value objectivec.IObject) int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("privateSetAppProperty:value:"), property, value)
	return rv
}
func (a AVAudioApplication) PrivateSetAppPropertyValueGuard(property objectivec.IObject, value objectivec.IObject, guard unsafe.Pointer) int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("privateSetAppProperty:value:guard:"), property, value, guard)
	return rv
}
func (a AVAudioApplication) PrivateSetInputMuteStateChangeHandlerError(handler func()) (bool, error) {
	_block0, _ := NewVoidBlock(handler)
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("privateSetInputMuteStateChangeHandler:error:"), _block0, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("privateSetInputMuteStateChangeHandler:error: returned NO with nil NSError")
	}
	return rv, nil

}
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
func (a AVAudioApplication) PrivateSetMXPropertyOnAllSessionsValue(sessions objectivec.IObject, value objectivec.IObject) int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("privateSetMXPropertyOnAllSessions:value:"), sessions, value)
	return rv
}
func (a AVAudioApplication) PrivateUpdateAppPropertyValueContext(property objectivec.IObject, value objectivec.IObject, context objectivec.IObject) int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("privateUpdateAppProperty:value:context:"), property, value, context)
	return rv
}
func (a AVAudioApplication) RequestRecordPermissionWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("requestRecordPermissionWithCompletionHandler:"), _block0)
}
func (a AVAudioApplication) SessionIDs() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("sessionIDs"))
	return objectivec.Object{ID: rv}
}
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
func (a AVAudioApplication) StemClickMutingEnabled() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("stemClickMutingEnabled"))
	return rv
}
func (a AVAudioApplication) InitDelegateForProcessProcessAttribution(process unsafe.Pointer, attribution objectivec.IObject) AVAudioApplication {
	rv := objc.SendIfResponds[AVAudioApplication](a.ID, objc.Sel("initDelegateForProcess:processAttribution:"), process, attribution)
	return rv
}
func (a AVAudioApplication) InitPrivate(private objectivec.IObject) AVAudioApplication {
	rv := objc.SendIfResponds[AVAudioApplication](a.ID, objc.Sel("initPrivate:"), private)
	return rv
}
func (a AVAudioApplication) InitProxyForProcess(process unsafe.Pointer) AVAudioApplication {
	rv := objc.SendIfResponds[AVAudioApplication](a.ID, objc.Sel("initProxyForProcess:"), process)
	return rv
}
func (a AVAudioApplication) InitWithSpecification(specification objectivec.IObject) AVAudioApplication {
	rv := objc.SendIfResponds[AVAudioApplication](a.ID, objc.Sel("initWithSpecification:"), specification)
	return rv
}

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
func (_AVAudioApplicationClass AVAudioApplicationClass) AppleTVSupportsEnhanceDialogue() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("appleTVSupportsEnhanceDialogue"))
	return rv
}
func (_AVAudioApplicationClass AVAudioApplicationClass) CurrentRouteSupportsEnhanceDialogue(dialogue []objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("currentRouteSupportsEnhanceDialogue:"), objectivec.IObjectSliceToNSArray(dialogue))
	return rv
}
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
func (_AVAudioApplicationClass AVAudioApplicationClass) IosDeviceSupportsEnhanceDialogue() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("iosDeviceSupportsEnhanceDialogue"))
	return rv
}
func (_AVAudioApplicationClass AVAudioApplicationClass) MuteRunningInputs(inputs []objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_AVAudioApplicationClass.class), objc.Sel("muteRunningInputs:"), objectivec.IObjectSliceToNSArray(inputs))
	return objectivec.Object{ID: rv}
}
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
func (_AVAudioApplicationClass AVAudioApplicationClass) ToggleInputMute(mute []objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("toggleInputMute:"), objectivec.IObjectSliceToNSArray(mute))
	return rv
}
func (_AVAudioApplicationClass AVAudioApplicationClass) VisionosDeviceSupportsEnhanceDialogue() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_AVAudioApplicationClass.class), objc.Sel("visionosDeviceSupportsEnhanceDialogue"))
	return rv
}

func (a AVAudioApplication) ClientID() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("clientID"))
	return rv
}
func (a AVAudioApplication) InputMuted() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("inputMuted"))
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
