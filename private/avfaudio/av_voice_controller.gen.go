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

// The class instance for the [AVVoiceController] class.
var (
	_AVVoiceControllerClass     AVVoiceControllerClass
	_AVVoiceControllerClassOnce sync.Once
)

func getAVVoiceControllerClass() AVVoiceControllerClass {
	_AVVoiceControllerClassOnce.Do(func() {
		_AVVoiceControllerClass = AVVoiceControllerClass{class: objc.GetClass("AVVoiceController")}
	})
	return _AVVoiceControllerClass
}

// GetAVVoiceControllerClass returns the class object for AVVoiceController.
func GetAVVoiceControllerClass() AVVoiceControllerClass {
	return getAVVoiceControllerClass()
}

type AVVoiceControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVoiceControllerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVoiceControllerClass) Alloc() AVVoiceController {
	rv := objc.Send[AVVoiceController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVoiceController.IsDeviceAvailableInLocalRouteError]
//   - [AVVoiceController._bringUpWithError]
//   - [AVVoiceController._teardownWithError]
//   - [AVVoiceController.ActivateAudioSessionForStreamIsPrewarmError]
//   - [AVVoiceController.ActivateAudioSessionForStreamIsPrewarmRecordModeError]
//   - [AVVoiceController.AlertPlaybackFinishedWithSettings]
//   - [AVVoiceController.AlertVolume]
//   - [AVVoiceController.SetAlertVolume]
//   - [AVVoiceController.BeganRecordingStatus]
//   - [AVVoiceController.CleanSlateWithError]
//   - [AVVoiceController.ConfigureAlertBehaviorForStreamCompletion]
//   - [AVVoiceController.ConfigureAlertBehaviorForStreamError]
//   - [AVVoiceController.ConfigureVoiceTriggerClientCompletionBlocks]
//   - [AVVoiceController.DeactivateAudioSessionForStreamWithOptionsCompletion]
//   - [AVVoiceController.DeactivateAudioSessionForStreamWithOptionsError]
//   - [AVVoiceController.DeactivateAudioSessionWithOptions]
//   - [AVVoiceController.EnableSmartRoutingConsiderationForStreamEnableError]
//   - [AVVoiceController.EnableTriangleModeForStreamEnableWithCompletion]
//   - [AVVoiceController.EncodeError]
//   - [AVVoiceController.EndpointDetectedAtTime]
//   - [AVVoiceController.EndpointerDelegate]
//   - [AVVoiceController.SetEndpointerDelegate]
//   - [AVVoiceController.FinishedRecordingStatus]
//   - [AVVoiceController.GetAveragePowerForStreamForChannel]
//   - [AVVoiceController.GetCurrentSessionState]
//   - [AVVoiceController.GetCurrentSessionStateForStream]
//   - [AVVoiceController.GetCurrentStreamState]
//   - [AVVoiceController.GetDeviceLatenciesForStreamWithCompletion]
//   - [AVVoiceController.GetInputChannelInfoForStreamCompletion]
//   - [AVVoiceController.GetPeakPowerForStreamForChannel]
//   - [AVVoiceController.GetPlaybackRouteForStreamWithCompletion]
//   - [AVVoiceController.GetPlaybackRouteForStreamWithError]
//   - [AVVoiceController.GetRecordBufferDurationForStream]
//   - [AVVoiceController.GetRecordDeviceInfoForStream]
//   - [AVVoiceController.GetRecordModeForStream]
//   - [AVVoiceController.GetRecordSettingsForStream]
//   - [AVVoiceController.HandleAudioHALDeviceWentAway]
//   - [AVVoiceController.HandlePluginDidPublishDeviceWithDevice]
//   - [AVVoiceController.HandlePluginDidUnpublishDeviceWithDevice]
//   - [AVVoiceController.Impl]
//   - [AVVoiceController.InterspeechPointDetectedAtTime]
//   - [AVVoiceController.IsDuckingSupportedOnPickedRouteForStreamError]
//   - [AVVoiceController.IsMeteringEnabledForStream]
//   - [AVVoiceController.Metrics]
//   - [AVVoiceController.MockPluginEndpoint]
//   - [AVVoiceController.NotifyEventOccuredError]
//   - [AVVoiceController.NotifyStreamInvalidated]
//   - [AVVoiceController.PlayAlertWithOverrideCompletion]
//   - [AVVoiceController.PlayAlertSoundForTypeOverrideMode]
//   - [AVVoiceController.PrepareRecordForStreamCompletion]
//   - [AVVoiceController.PrepareRecordForStreamError]
//   - [AVVoiceController.RecordDelegate]
//   - [AVVoiceController.SetRecordDelegate]
//   - [AVVoiceController.RecordEndWaitTime]
//   - [AVVoiceController.SetRecordEndWaitTime]
//   - [AVVoiceController.RecordEndpointMode]
//   - [AVVoiceController.SetRecordEndpointMode]
//   - [AVVoiceController.RecordInterspeechWaitTime]
//   - [AVVoiceController.SetRecordInterspeechWaitTime]
//   - [AVVoiceController.RecordStartWaitTime]
//   - [AVVoiceController.SetRecordStartWaitTime]
//   - [AVVoiceController.RemoveStreamCompletion]
//   - [AVVoiceController.SetAlertSoundFromURLForType]
//   - [AVVoiceController.SetAnnounceCallsEnabledForStreamEnable]
//   - [AVVoiceController.SetContextCompletion]
//   - [AVVoiceController.SetContextError]
//   - [AVVoiceController.SetContextStreamTypeError]
//   - [AVVoiceController.SetContextForStreamForStreamError]
//   - [AVVoiceController.SetDuckOthersForStreamWithSettingsError]
//   - [AVVoiceController.SetEnableInterruptionByRecordingClientsForStreamEnableError]
//   - [AVVoiceController.SetRecordModeForStreamRecordModeError]
//   - [AVVoiceController.SetRecordStatusChangeBlock]
//   - [AVVoiceController.StartKeepAliveQueueForStreamCompletion]
//   - [AVVoiceController.StartRecordForStreamCompletion]
//   - [AVVoiceController.StartRecordForStreamError]
//   - [AVVoiceController.StartRecordWithSettingsCompletionAlertCompletionAudioCallback]
//   - [AVVoiceController.StartpointDetected]
//   - [AVVoiceController.StopKeepAliveQueueForStreamCompletion]
//   - [AVVoiceController.StopRecordForStreamCompletion]
//   - [AVVoiceController.StopRecordForStreamError]
//   - [AVVoiceController.TeardownWithError]
//   - [AVVoiceController.UpdateMeterForStream]
//   - [AVVoiceController.InitVoiceControllerForClientWithError]
//   - [AVVoiceController.InitWithError]
//   - [AVVoiceController.DebugDescription]
//   - [AVVoiceController.Description]
//   - [AVVoiceController.Hash]
//   - [AVVoiceController.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController
type AVVoiceController struct {
	objectivec.Object
}

// AVVoiceControllerFromID constructs a [AVVoiceController] from an objc.ID.
func AVVoiceControllerFromID(id objc.ID) AVVoiceController {
	return AVVoiceController{objectivec.Object{ID: id}}
}

// Ensure AVVoiceController implements IAVVoiceController.
var _ IAVVoiceController = AVVoiceController{}

// An interface definition for the [AVVoiceController] class.
//
// # Methods
//
//   - [IAVVoiceController.IsDeviceAvailableInLocalRouteError]
//   - [IAVVoiceController._bringUpWithError]
//   - [IAVVoiceController._teardownWithError]
//   - [IAVVoiceController.ActivateAudioSessionForStreamIsPrewarmError]
//   - [IAVVoiceController.ActivateAudioSessionForStreamIsPrewarmRecordModeError]
//   - [IAVVoiceController.AlertPlaybackFinishedWithSettings]
//   - [IAVVoiceController.AlertVolume]
//   - [IAVVoiceController.SetAlertVolume]
//   - [IAVVoiceController.BeganRecordingStatus]
//   - [IAVVoiceController.CleanSlateWithError]
//   - [IAVVoiceController.ConfigureAlertBehaviorForStreamCompletion]
//   - [IAVVoiceController.ConfigureAlertBehaviorForStreamError]
//   - [IAVVoiceController.ConfigureVoiceTriggerClientCompletionBlocks]
//   - [IAVVoiceController.DeactivateAudioSessionForStreamWithOptionsCompletion]
//   - [IAVVoiceController.DeactivateAudioSessionForStreamWithOptionsError]
//   - [IAVVoiceController.DeactivateAudioSessionWithOptions]
//   - [IAVVoiceController.EnableSmartRoutingConsiderationForStreamEnableError]
//   - [IAVVoiceController.EnableTriangleModeForStreamEnableWithCompletion]
//   - [IAVVoiceController.EncodeError]
//   - [IAVVoiceController.EndpointDetectedAtTime]
//   - [IAVVoiceController.EndpointerDelegate]
//   - [IAVVoiceController.SetEndpointerDelegate]
//   - [IAVVoiceController.FinishedRecordingStatus]
//   - [IAVVoiceController.GetAveragePowerForStreamForChannel]
//   - [IAVVoiceController.GetCurrentSessionState]
//   - [IAVVoiceController.GetCurrentSessionStateForStream]
//   - [IAVVoiceController.GetCurrentStreamState]
//   - [IAVVoiceController.GetDeviceLatenciesForStreamWithCompletion]
//   - [IAVVoiceController.GetInputChannelInfoForStreamCompletion]
//   - [IAVVoiceController.GetPeakPowerForStreamForChannel]
//   - [IAVVoiceController.GetPlaybackRouteForStreamWithCompletion]
//   - [IAVVoiceController.GetPlaybackRouteForStreamWithError]
//   - [IAVVoiceController.GetRecordBufferDurationForStream]
//   - [IAVVoiceController.GetRecordDeviceInfoForStream]
//   - [IAVVoiceController.GetRecordModeForStream]
//   - [IAVVoiceController.GetRecordSettingsForStream]
//   - [IAVVoiceController.HandleAudioHALDeviceWentAway]
//   - [IAVVoiceController.HandlePluginDidPublishDeviceWithDevice]
//   - [IAVVoiceController.HandlePluginDidUnpublishDeviceWithDevice]
//   - [IAVVoiceController.Impl]
//   - [IAVVoiceController.InterspeechPointDetectedAtTime]
//   - [IAVVoiceController.IsDuckingSupportedOnPickedRouteForStreamError]
//   - [IAVVoiceController.IsMeteringEnabledForStream]
//   - [IAVVoiceController.Metrics]
//   - [IAVVoiceController.MockPluginEndpoint]
//   - [IAVVoiceController.NotifyEventOccuredError]
//   - [IAVVoiceController.NotifyStreamInvalidated]
//   - [IAVVoiceController.PlayAlertWithOverrideCompletion]
//   - [IAVVoiceController.PlayAlertSoundForTypeOverrideMode]
//   - [IAVVoiceController.PrepareRecordForStreamCompletion]
//   - [IAVVoiceController.PrepareRecordForStreamError]
//   - [IAVVoiceController.RecordDelegate]
//   - [IAVVoiceController.SetRecordDelegate]
//   - [IAVVoiceController.RecordEndWaitTime]
//   - [IAVVoiceController.SetRecordEndWaitTime]
//   - [IAVVoiceController.RecordEndpointMode]
//   - [IAVVoiceController.SetRecordEndpointMode]
//   - [IAVVoiceController.RecordInterspeechWaitTime]
//   - [IAVVoiceController.SetRecordInterspeechWaitTime]
//   - [IAVVoiceController.RecordStartWaitTime]
//   - [IAVVoiceController.SetRecordStartWaitTime]
//   - [IAVVoiceController.RemoveStreamCompletion]
//   - [IAVVoiceController.SetAlertSoundFromURLForType]
//   - [IAVVoiceController.SetAnnounceCallsEnabledForStreamEnable]
//   - [IAVVoiceController.SetContextCompletion]
//   - [IAVVoiceController.SetContextError]
//   - [IAVVoiceController.SetContextStreamTypeError]
//   - [IAVVoiceController.SetContextForStreamForStreamError]
//   - [IAVVoiceController.SetDuckOthersForStreamWithSettingsError]
//   - [IAVVoiceController.SetEnableInterruptionByRecordingClientsForStreamEnableError]
//   - [IAVVoiceController.SetRecordModeForStreamRecordModeError]
//   - [IAVVoiceController.SetRecordStatusChangeBlock]
//   - [IAVVoiceController.StartKeepAliveQueueForStreamCompletion]
//   - [IAVVoiceController.StartRecordForStreamCompletion]
//   - [IAVVoiceController.StartRecordForStreamError]
//   - [IAVVoiceController.StartRecordWithSettingsCompletionAlertCompletionAudioCallback]
//   - [IAVVoiceController.StartpointDetected]
//   - [IAVVoiceController.StopKeepAliveQueueForStreamCompletion]
//   - [IAVVoiceController.StopRecordForStreamCompletion]
//   - [IAVVoiceController.StopRecordForStreamError]
//   - [IAVVoiceController.TeardownWithError]
//   - [IAVVoiceController.UpdateMeterForStream]
//   - [IAVVoiceController.InitVoiceControllerForClientWithError]
//   - [IAVVoiceController.InitWithError]
//   - [IAVVoiceController.DebugDescription]
//   - [IAVVoiceController.Description]
//   - [IAVVoiceController.Hash]
//   - [IAVVoiceController.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController
type IAVVoiceController interface {
	objectivec.IObject

	// Topic: Methods

	IsDeviceAvailableInLocalRouteError(route objectivec.IObject) (bool, error)
	_bringUpWithError(up int64) (int64, error)
	_teardownWithError() error
	ActivateAudioSessionForStreamIsPrewarmError(stream uint64, prewarm bool) (bool, error)
	ActivateAudioSessionForStreamIsPrewarmRecordModeError(stream uint64, prewarm bool, mode bool) (bool, error)
	AlertPlaybackFinishedWithSettings(settings objectivec.IObject)
	AlertVolume() float32
	SetAlertVolume(value float32)
	BeganRecordingStatus(recording uint64, status int)
	CleanSlateWithError() error
	ConfigureAlertBehaviorForStreamCompletion(stream objectivec.IObject, completion VoidHandler)
	ConfigureAlertBehaviorForStreamError(stream objectivec.IObject) (bool, error)
	ConfigureVoiceTriggerClientCompletionBlocks()
	DeactivateAudioSessionForStreamWithOptionsCompletion(stream uint64, options uint64, completion VoidHandler)
	DeactivateAudioSessionForStreamWithOptionsError(stream uint64, options uint64) error
	DeactivateAudioSessionWithOptions(options uint64)
	EnableSmartRoutingConsiderationForStreamEnableError(stream uint64, enable bool) (bool, error)
	EnableTriangleModeForStreamEnableWithCompletion(stream uint64, enable bool, completion VoidHandler)
	EncodeError(error_ int)
	EndpointDetectedAtTime(time float64)
	EndpointerDelegate() objectivec.IObject
	SetEndpointerDelegate(value objectivec.IObject)
	FinishedRecordingStatus(recording uint64, status int)
	GetAveragePowerForStreamForChannel(stream uint64, channel uint64) float32
	GetCurrentSessionState() int64
	GetCurrentSessionStateForStream(stream uint64) int64
	GetCurrentStreamState(state uint64) int64
	GetDeviceLatenciesForStreamWithCompletion(stream uint64, completion VoidHandler)
	GetInputChannelInfoForStreamCompletion(stream uint64, completion VoidHandler)
	GetPeakPowerForStreamForChannel(stream uint64, channel uint64) float32
	GetPlaybackRouteForStreamWithCompletion(stream uint64, completion VoidHandler)
	GetPlaybackRouteForStreamWithError(stream uint64) (objectivec.IObject, error)
	GetRecordBufferDurationForStream(stream uint64) float64
	GetRecordDeviceInfoForStream(stream uint64) objectivec.IObject
	GetRecordModeForStream(stream uint64) int64
	GetRecordSettingsForStream(stream uint64) objectivec.IObject
	HandleAudioHALDeviceWentAway(away uint64)
	HandlePluginDidPublishDeviceWithDevice(device objectivec.IObject, device2 objectivec.IObject)
	HandlePluginDidUnpublishDeviceWithDevice(device objectivec.IObject, device2 objectivec.IObject)
	Impl() objectivec.IObject
	InterspeechPointDetectedAtTime(time float64)
	IsDuckingSupportedOnPickedRouteForStreamError(stream uint64) (bool, error)
	IsMeteringEnabledForStream(stream uint64) bool
	Metrics() foundation.INSDictionary
	MockPluginEndpoint() objectivec.IObject
	NotifyEventOccuredError(occured uint64, error_ objectivec.IObject)
	NotifyStreamInvalidated(invalidated uint64)
	PlayAlertWithOverrideCompletion(alert int, override int64, completion VoidHandler)
	PlayAlertSoundForTypeOverrideMode(type_ int, mode int64) bool
	PrepareRecordForStreamCompletion(stream objectivec.IObject, completion VoidHandler)
	PrepareRecordForStreamError(stream objectivec.IObject) (bool, error)
	RecordDelegate() objectivec.IObject
	SetRecordDelegate(value objectivec.IObject)
	RecordEndWaitTime() float64
	SetRecordEndWaitTime(value float64)
	RecordEndpointMode() int
	SetRecordEndpointMode(value int)
	RecordInterspeechWaitTime() float64
	SetRecordInterspeechWaitTime(value float64)
	RecordStartWaitTime() float64
	SetRecordStartWaitTime(value float64)
	RemoveStreamCompletion(stream uint64, completion VoidHandler)
	SetAlertSoundFromURLForType(url foundation.INSURL, type_ int) bool
	SetAnnounceCallsEnabledForStreamEnable(stream uint64, enable bool) bool
	SetContextCompletion(context objectivec.IObject, completion VoidHandler)
	SetContextError(context objectivec.IObject) (uint64, error)
	SetContextStreamTypeError(context objectivec.IObject, type_ unsafe.Pointer) (uint64, error)
	SetContextForStreamForStreamError(stream objectivec.IObject, stream2 uint64) (bool, error)
	SetDuckOthersForStreamWithSettingsError(stream uint64, settings objectivec.IObject) (bool, error)
	SetEnableInterruptionByRecordingClientsForStreamEnableError(stream uint64, enable bool) (bool, error)
	SetRecordModeForStreamRecordModeError(stream uint64, mode int64) (bool, error)
	SetRecordStatusChangeBlock(block VoidHandler)
	StartKeepAliveQueueForStreamCompletion(stream uint64, completion VoidHandler)
	StartRecordForStreamCompletion(stream objectivec.IObject, completion VoidHandler)
	StartRecordForStreamError(stream objectivec.IObject) (bool, error)
	StartRecordWithSettingsCompletionAlertCompletionAudioCallback(settings objectivec.IObject, completion VoidHandler, completion2 VoidHandler, callback VoidHandler)
	StartpointDetected()
	StopKeepAliveQueueForStreamCompletion(stream uint64, completion VoidHandler)
	StopRecordForStreamCompletion(stream uint64, completion VoidHandler)
	StopRecordForStreamError(stream uint64) (bool, error)
	TeardownWithError() error
	UpdateMeterForStream(stream uint64) bool
	InitVoiceControllerForClientWithError(client int64) (AVVoiceController, error)
	InitWithError() (AVVoiceController, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v AVVoiceController) Init() AVVoiceController {
	rv := objc.Send[AVVoiceController](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVoiceController) Autorelease() AVVoiceController {
	rv := objc.Send[AVVoiceController](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVoiceController creates a new AVVoiceController instance.
func NewAVVoiceController() AVVoiceController {
	class := getAVVoiceControllerClass()
	rv := objc.Send[AVVoiceController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/initVoiceControllerForClient:withError:
func NewVoiceControllerVoiceControllerForClientWithError(client int64) (AVVoiceController, error) {
	var errorPtr objc.ID
	instance := getAVVoiceControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initVoiceControllerForClient:withError:"), client, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVVoiceController{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVVoiceControllerFromID(rv), nil
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/initWithError:
func NewVoiceControllerWithError() (AVVoiceController, error) {
	var errorPtr objc.ID
	instance := getAVVoiceControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVVoiceController{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVVoiceControllerFromID(rv), nil
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/IsDeviceAvailableInLocalRoute:error:
func (v AVVoiceController) IsDeviceAvailableInLocalRouteError(route objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("IsDeviceAvailableInLocalRoute:error:"), route, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("IsDeviceAvailableInLocalRoute:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/_bringUp:withError:
func (v AVVoiceController) _bringUpWithError(up int64) (int64, error) {
	var errorPtr objc.ID
	rv := objc.Send[int64](v.ID, objc.Sel("_bringUp:withError:"), up, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// BringUpWithError is an exported wrapper for the private method _bringUpWithError.
func (v AVVoiceController) BringUpWithError(up int64) (int64, error) {
	return v._bringUpWithError(up)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/_teardownWithError:
func (v AVVoiceController) _teardownWithError() error {
	var errorPtr objc.ID
	objc.Send[struct{}](v.ID, objc.Sel("_teardownWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/activateAudioSessionForStream:isPrewarm:error:
func (v AVVoiceController) ActivateAudioSessionForStreamIsPrewarmError(stream uint64, prewarm bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("activateAudioSessionForStream:isPrewarm:error:"), stream, prewarm, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("activateAudioSessionForStream:isPrewarm:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/activateAudioSessionForStream:isPrewarm:recordMode:error:
func (v AVVoiceController) ActivateAudioSessionForStreamIsPrewarmRecordModeError(stream uint64, prewarm bool, mode bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("activateAudioSessionForStream:isPrewarm:recordMode:error:"), stream, prewarm, mode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("activateAudioSessionForStream:isPrewarm:recordMode:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/alertPlaybackFinishedWithSettings:
func (v AVVoiceController) AlertPlaybackFinishedWithSettings(settings objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("alertPlaybackFinishedWithSettings:"), settings)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/beganRecording:status:
func (v AVVoiceController) BeganRecordingStatus(recording uint64, status int) {
	objc.Send[objc.ID](v.ID, objc.Sel("beganRecording:status:"), recording, status)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/cleanSlateWithError:
func (v AVVoiceController) CleanSlateWithError() error {
	var errorPtr objc.ID
	objc.Send[struct{}](v.ID, objc.Sel("cleanSlateWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/configureAlertBehaviorForStream:completion:
func (v AVVoiceController) ConfigureAlertBehaviorForStreamCompletion(stream objectivec.IObject, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("configureAlertBehaviorForStream:completion:"), stream, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/configureAlertBehaviorForStream:error:
func (v AVVoiceController) ConfigureAlertBehaviorForStreamError(stream objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("configureAlertBehaviorForStream:error:"), stream, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("configureAlertBehaviorForStream:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/configureVoiceTriggerClientCompletionBlocks
func (v AVVoiceController) ConfigureVoiceTriggerClientCompletionBlocks() {
	objc.Send[objc.ID](v.ID, objc.Sel("configureVoiceTriggerClientCompletionBlocks"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/deactivateAudioSessionForStream:withOptions:completion:
func (v AVVoiceController) DeactivateAudioSessionForStreamWithOptionsCompletion(stream uint64, options uint64, completion VoidHandler) {
	_block2, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("deactivateAudioSessionForStream:withOptions:completion:"), stream, options, _block2)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/deactivateAudioSessionForStream:withOptions:error:
func (v AVVoiceController) DeactivateAudioSessionForStreamWithOptionsError(stream uint64, options uint64) error {
	var errorPtr objc.ID
	objc.Send[struct{}](v.ID, objc.Sel("deactivateAudioSessionForStream:withOptions:error:"), stream, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/deactivateAudioSessionWithOptions:
func (v AVVoiceController) DeactivateAudioSessionWithOptions(options uint64) {
	objc.Send[objc.ID](v.ID, objc.Sel("deactivateAudioSessionWithOptions:"), options)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/enableSmartRoutingConsiderationForStream:enable:error:
func (v AVVoiceController) EnableSmartRoutingConsiderationForStreamEnableError(stream uint64, enable bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("enableSmartRoutingConsiderationForStream:enable:error:"), stream, enable, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enableSmartRoutingConsiderationForStream:enable:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/enableTriangleModeForStream:enable:withCompletion:
func (v AVVoiceController) EnableTriangleModeForStreamEnableWithCompletion(stream uint64, enable bool, completion VoidHandler) {
	_block2, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("enableTriangleModeForStream:enable:withCompletion:"), stream, enable, _block2)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/encodeError:
func (v AVVoiceController) EncodeError(error_ int) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeError:"), error_)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/endpointDetectedAtTime:
func (v AVVoiceController) EndpointDetectedAtTime(time float64) {
	objc.Send[objc.ID](v.ID, objc.Sel("endpointDetectedAtTime:"), time)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/finishedRecording:status:
func (v AVVoiceController) FinishedRecordingStatus(recording uint64, status int) {
	objc.Send[objc.ID](v.ID, objc.Sel("finishedRecording:status:"), recording, status)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getAveragePowerForStream:forChannel:
func (v AVVoiceController) GetAveragePowerForStreamForChannel(stream uint64, channel uint64) float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("getAveragePowerForStream:forChannel:"), stream, channel)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getCurrentSessionState
func (v AVVoiceController) GetCurrentSessionState() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("getCurrentSessionState"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getCurrentSessionStateForStream:
func (v AVVoiceController) GetCurrentSessionStateForStream(stream uint64) int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("getCurrentSessionStateForStream:"), stream)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getCurrentStreamState:
func (v AVVoiceController) GetCurrentStreamState(state uint64) int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("getCurrentStreamState:"), state)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getDeviceLatenciesForStream:withCompletion:
func (v AVVoiceController) GetDeviceLatenciesForStreamWithCompletion(stream uint64, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("getDeviceLatenciesForStream:withCompletion:"), stream, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getInputChannelInfoForStream:completion:
func (v AVVoiceController) GetInputChannelInfoForStreamCompletion(stream uint64, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("getInputChannelInfoForStream:completion:"), stream, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getPeakPowerForStream:forChannel:
func (v AVVoiceController) GetPeakPowerForStreamForChannel(stream uint64, channel uint64) float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("getPeakPowerForStream:forChannel:"), stream, channel)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getPlaybackRouteForStream:withCompletion:
func (v AVVoiceController) GetPlaybackRouteForStreamWithCompletion(stream uint64, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("getPlaybackRouteForStream:withCompletion:"), stream, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getPlaybackRouteForStream:withError:
func (v AVVoiceController) GetPlaybackRouteForStreamWithError(stream uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("getPlaybackRouteForStream:withError:"), stream, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getRecordBufferDurationForStream:
func (v AVVoiceController) GetRecordBufferDurationForStream(stream uint64) float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("getRecordBufferDurationForStream:"), stream)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getRecordDeviceInfoForStream:
func (v AVVoiceController) GetRecordDeviceInfoForStream(stream uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("getRecordDeviceInfoForStream:"), stream)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getRecordModeForStream:
func (v AVVoiceController) GetRecordModeForStream(stream uint64) int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("getRecordModeForStream:"), stream)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/getRecordSettingsForStream:
func (v AVVoiceController) GetRecordSettingsForStream(stream uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("getRecordSettingsForStream:"), stream)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/handleAudioHALDeviceWentAway:
func (v AVVoiceController) HandleAudioHALDeviceWentAway(away uint64) {
	objc.Send[objc.ID](v.ID, objc.Sel("handleAudioHALDeviceWentAway:"), away)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/handlePluginDidPublishDevice:withDevice:
func (v AVVoiceController) HandlePluginDidPublishDeviceWithDevice(device objectivec.IObject, device2 objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("handlePluginDidPublishDevice:withDevice:"), device, device2)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/handlePluginDidUnpublishDevice:withDevice:
func (v AVVoiceController) HandlePluginDidUnpublishDeviceWithDevice(device objectivec.IObject, device2 objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("handlePluginDidUnpublishDevice:withDevice:"), device, device2)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/impl
func (v AVVoiceController) Impl() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("impl"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/interspeechPointDetectedAtTime:
func (v AVVoiceController) InterspeechPointDetectedAtTime(time float64) {
	objc.Send[objc.ID](v.ID, objc.Sel("interspeechPointDetectedAtTime:"), time)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/isDuckingSupportedOnPickedRouteForStream:error:
func (v AVVoiceController) IsDuckingSupportedOnPickedRouteForStreamError(stream uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("isDuckingSupportedOnPickedRouteForStream:error:"), stream, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("isDuckingSupportedOnPickedRouteForStream:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/isMeteringEnabledForStream:
func (v AVVoiceController) IsMeteringEnabledForStream(stream uint64) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isMeteringEnabledForStream:"), stream)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/mockPluginEndpoint
func (v AVVoiceController) MockPluginEndpoint() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("mockPluginEndpoint"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/notifyEventOccured:error:
func (v AVVoiceController) NotifyEventOccuredError(occured uint64, error_ objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("notifyEventOccured:error:"), occured, error_)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/notifyStreamInvalidated:
func (v AVVoiceController) NotifyStreamInvalidated(invalidated uint64) {
	objc.Send[objc.ID](v.ID, objc.Sel("notifyStreamInvalidated:"), invalidated)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/playAlert:withOverride:completion:
func (v AVVoiceController) PlayAlertWithOverrideCompletion(alert int, override int64, completion VoidHandler) {
	_block2, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("playAlert:withOverride:completion:"), alert, override, _block2)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/playAlertSoundForType:overrideMode:
func (v AVVoiceController) PlayAlertSoundForTypeOverrideMode(type_ int, mode int64) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("playAlertSoundForType:overrideMode:"), type_, mode)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/prepareRecordForStream:completion:
func (v AVVoiceController) PrepareRecordForStreamCompletion(stream objectivec.IObject, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("prepareRecordForStream:completion:"), stream, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/prepareRecordForStream:error:
func (v AVVoiceController) PrepareRecordForStreamError(stream objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("prepareRecordForStream:error:"), stream, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareRecordForStream:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/removeStream:completion:
func (v AVVoiceController) RemoveStreamCompletion(stream uint64, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("removeStream:completion:"), stream, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/setAlertSoundFromURL:forType:
func (v AVVoiceController) SetAlertSoundFromURLForType(url foundation.INSURL, type_ int) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("setAlertSoundFromURL:forType:"), url, type_)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/setAnnounceCallsEnabledForStream:enable:
func (v AVVoiceController) SetAnnounceCallsEnabledForStreamEnable(stream uint64, enable bool) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("setAnnounceCallsEnabledForStream:enable:"), stream, enable)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/setContext:completion:
func (v AVVoiceController) SetContextCompletion(context objectivec.IObject, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("setContext:completion:"), context, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/setContext:error:
func (v AVVoiceController) SetContextError(context objectivec.IObject) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](v.ID, objc.Sel("setContext:error:"), context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/setContext:streamType:error:
func (v AVVoiceController) SetContextStreamTypeError(context objectivec.IObject, type_ unsafe.Pointer) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](v.ID, objc.Sel("setContext:streamType:error:"), context, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/setContextForStream:forStream:error:
func (v AVVoiceController) SetContextForStreamForStreamError(stream objectivec.IObject, stream2 uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("setContextForStream:forStream:error:"), stream, stream2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setContextForStream:forStream:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/setDuckOthersForStream:withSettings:error:
func (v AVVoiceController) SetDuckOthersForStreamWithSettingsError(stream uint64, settings objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("setDuckOthersForStream:withSettings:error:"), stream, settings, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setDuckOthersForStream:withSettings:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/setEnableInterruptionByRecordingClientsForStream:enable:error:
func (v AVVoiceController) SetEnableInterruptionByRecordingClientsForStreamEnableError(stream uint64, enable bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("setEnableInterruptionByRecordingClientsForStream:enable:error:"), stream, enable, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setEnableInterruptionByRecordingClientsForStream:enable:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/setRecordModeForStream:recordMode:error:
func (v AVVoiceController) SetRecordModeForStreamRecordModeError(stream uint64, mode int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("setRecordModeForStream:recordMode:error:"), stream, mode, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setRecordModeForStream:recordMode:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/setRecordStatusChangeBlock:
func (v AVVoiceController) SetRecordStatusChangeBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setRecordStatusChangeBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/startKeepAliveQueueForStream:completion:
func (v AVVoiceController) StartKeepAliveQueueForStreamCompletion(stream uint64, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("startKeepAliveQueueForStream:completion:"), stream, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/startRecordForStream:completion:
func (v AVVoiceController) StartRecordForStreamCompletion(stream objectivec.IObject, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("startRecordForStream:completion:"), stream, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/startRecordForStream:error:
func (v AVVoiceController) StartRecordForStreamError(stream objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("startRecordForStream:error:"), stream, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startRecordForStream:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/startRecordWithSettings:completion:alertCompletion:audioCallback:
func (v AVVoiceController) StartRecordWithSettingsCompletionAlertCompletionAudioCallback(settings objectivec.IObject, completion VoidHandler, completion2 VoidHandler, callback VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	_block2, _ := NewVoidBlock(completion2)
	_block3, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](v.ID, objc.Sel("startRecordWithSettings:completion:alertCompletion:audioCallback:"), settings, _block1, _block2, _block3)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/startpointDetected
func (v AVVoiceController) StartpointDetected() {
	objc.Send[objc.ID](v.ID, objc.Sel("startpointDetected"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/stopKeepAliveQueueForStream:completion:
func (v AVVoiceController) StopKeepAliveQueueForStreamCompletion(stream uint64, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("stopKeepAliveQueueForStream:completion:"), stream, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/stopRecordForStream:completion:
func (v AVVoiceController) StopRecordForStreamCompletion(stream uint64, completion VoidHandler) {
	_block1, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("stopRecordForStream:completion:"), stream, _block1)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/stopRecordForStream:error:
func (v AVVoiceController) StopRecordForStreamError(stream uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("stopRecordForStream:error:"), stream, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("stopRecordForStream:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/teardownWithError:
func (v AVVoiceController) TeardownWithError() error {
	var errorPtr objc.ID
	objc.Send[struct{}](v.ID, objc.Sel("teardownWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/updateMeterForStream:
func (v AVVoiceController) UpdateMeterForStream(stream uint64) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("updateMeterForStream:"), stream)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/initVoiceControllerForClient:withError:
func (v AVVoiceController) InitVoiceControllerForClientWithError(client int64) (AVVoiceController, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initVoiceControllerForClient:withError:"), client, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVVoiceController{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVVoiceControllerFromID(rv), nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/initWithError:
func (v AVVoiceController) InitWithError() (AVVoiceController, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVVoiceController{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVVoiceControllerFromID(rv), nil

}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/alertVolume
func (v AVVoiceController) AlertVolume() float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("alertVolume"))
	return rv
}
func (v AVVoiceController) SetAlertVolume(value float32) {
	objc.Send[struct{}](v.ID, objc.Sel("setAlertVolume:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/debugDescription
func (v AVVoiceController) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/description
func (v AVVoiceController) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/endpointerDelegate
func (v AVVoiceController) EndpointerDelegate() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("endpointerDelegate"))
	return objectivec.Object{ID: rv}
}
func (v AVVoiceController) SetEndpointerDelegate(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setEndpointerDelegate:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/hash
func (v AVVoiceController) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/metrics
func (v AVVoiceController) Metrics() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("metrics"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/recordDelegate
func (v AVVoiceController) RecordDelegate() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("recordDelegate"))
	return objectivec.Object{ID: rv}
}
func (v AVVoiceController) SetRecordDelegate(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setRecordDelegate:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/recordEndWaitTime
func (v AVVoiceController) RecordEndWaitTime() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("recordEndWaitTime"))
	return rv
}
func (v AVVoiceController) SetRecordEndWaitTime(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setRecordEndWaitTime:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/recordEndpointMode
func (v AVVoiceController) RecordEndpointMode() int {
	rv := objc.Send[int](v.ID, objc.Sel("recordEndpointMode"))
	return rv
}
func (v AVVoiceController) SetRecordEndpointMode(value int) {
	objc.Send[struct{}](v.ID, objc.Sel("setRecordEndpointMode:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/recordInterspeechWaitTime
func (v AVVoiceController) RecordInterspeechWaitTime() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("recordInterspeechWaitTime"))
	return rv
}
func (v AVVoiceController) SetRecordInterspeechWaitTime(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setRecordInterspeechWaitTime:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/recordStartWaitTime
func (v AVVoiceController) RecordStartWaitTime() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("recordStartWaitTime"))
	return rv
}
func (v AVVoiceController) SetRecordStartWaitTime(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setRecordStartWaitTime:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceController/superclass
func (v AVVoiceController) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

// ConfigureAlertBehaviorForStreamCompletionSync is a synchronous wrapper around [AVVoiceController.ConfigureAlertBehaviorForStreamCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) ConfigureAlertBehaviorForStreamCompletionSync(ctx context.Context, stream objectivec.IObject) error {
	done := make(chan struct{}, 1)
	v.ConfigureAlertBehaviorForStreamCompletion(stream, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// DeactivateAudioSessionForStreamWithOptionsCompletionSync is a synchronous wrapper around [AVVoiceController.DeactivateAudioSessionForStreamWithOptionsCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) DeactivateAudioSessionForStreamWithOptionsCompletionSync(ctx context.Context, stream uint64, options uint64) error {
	done := make(chan struct{}, 1)
	v.DeactivateAudioSessionForStreamWithOptionsCompletion(stream, options, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnableTriangleModeForStreamEnable is a synchronous wrapper around [AVVoiceController.EnableTriangleModeForStreamEnableWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) EnableTriangleModeForStreamEnable(ctx context.Context, stream uint64, enable bool) error {
	done := make(chan struct{}, 1)
	v.EnableTriangleModeForStreamEnableWithCompletion(stream, enable, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetDeviceLatenciesForStream is a synchronous wrapper around [AVVoiceController.GetDeviceLatenciesForStreamWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) GetDeviceLatenciesForStream(ctx context.Context, stream uint64) error {
	done := make(chan struct{}, 1)
	v.GetDeviceLatenciesForStreamWithCompletion(stream, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetInputChannelInfoForStreamCompletionSync is a synchronous wrapper around [AVVoiceController.GetInputChannelInfoForStreamCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) GetInputChannelInfoForStreamCompletionSync(ctx context.Context, stream uint64) error {
	done := make(chan struct{}, 1)
	v.GetInputChannelInfoForStreamCompletion(stream, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetPlaybackRouteForStream is a synchronous wrapper around [AVVoiceController.GetPlaybackRouteForStreamWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) GetPlaybackRouteForStream(ctx context.Context, stream uint64) error {
	done := make(chan struct{}, 1)
	v.GetPlaybackRouteForStreamWithCompletion(stream, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PlayAlertWithOverrideCompletionSync is a synchronous wrapper around [AVVoiceController.PlayAlertWithOverrideCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) PlayAlertWithOverrideCompletionSync(ctx context.Context, alert int, override int64) error {
	done := make(chan struct{}, 1)
	v.PlayAlertWithOverrideCompletion(alert, override, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PrepareRecordForStreamCompletionSync is a synchronous wrapper around [AVVoiceController.PrepareRecordForStreamCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) PrepareRecordForStreamCompletionSync(ctx context.Context, stream objectivec.IObject) error {
	done := make(chan struct{}, 1)
	v.PrepareRecordForStreamCompletion(stream, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RemoveStreamCompletionSync is a synchronous wrapper around [AVVoiceController.RemoveStreamCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) RemoveStreamCompletionSync(ctx context.Context, stream uint64) error {
	done := make(chan struct{}, 1)
	v.RemoveStreamCompletion(stream, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetContextCompletionSync is a synchronous wrapper around [AVVoiceController.SetContextCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) SetContextCompletionSync(ctx context.Context, context objectivec.IObject) error {
	done := make(chan struct{}, 1)
	v.SetContextCompletion(context, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetRecordStatusChangeBlockSync is a synchronous wrapper around [AVVoiceController.SetRecordStatusChangeBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) SetRecordStatusChangeBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetRecordStatusChangeBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StartKeepAliveQueueForStreamCompletionSync is a synchronous wrapper around [AVVoiceController.StartKeepAliveQueueForStreamCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) StartKeepAliveQueueForStreamCompletionSync(ctx context.Context, stream uint64) error {
	done := make(chan struct{}, 1)
	v.StartKeepAliveQueueForStreamCompletion(stream, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StartRecordForStreamCompletionSync is a synchronous wrapper around [AVVoiceController.StartRecordForStreamCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) StartRecordForStreamCompletionSync(ctx context.Context, stream objectivec.IObject) error {
	done := make(chan struct{}, 1)
	v.StartRecordForStreamCompletion(stream, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StartRecordWithSettingsCompletionAlertCompletionAudioCallbackSync is a synchronous wrapper around [AVVoiceController.StartRecordWithSettingsCompletionAlertCompletionAudioCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) StartRecordWithSettingsCompletionAlertCompletionAudioCallbackSync(ctx context.Context, settings objectivec.IObject, completion VoidHandler, completion2 VoidHandler) error {
	done := make(chan struct{}, 1)
	v.StartRecordWithSettingsCompletionAlertCompletionAudioCallback(settings, completion, completion2, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StopKeepAliveQueueForStreamCompletionSync is a synchronous wrapper around [AVVoiceController.StopKeepAliveQueueForStreamCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) StopKeepAliveQueueForStreamCompletionSync(ctx context.Context, stream uint64) error {
	done := make(chan struct{}, 1)
	v.StopKeepAliveQueueForStreamCompletion(stream, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StopRecordForStreamCompletionSync is a synchronous wrapper around [AVVoiceController.StopRecordForStreamCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceController) StopRecordForStreamCompletionSync(ctx context.Context, stream uint64) error {
	done := make(chan struct{}, 1)
	v.StopRecordForStreamCompletion(stream, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
