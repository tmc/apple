// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVoiceTriggerClient] class.
var (
	_AVVoiceTriggerClientClass     AVVoiceTriggerClientClass
	_AVVoiceTriggerClientClassOnce sync.Once
)

func getAVVoiceTriggerClientClass() AVVoiceTriggerClientClass {
	_AVVoiceTriggerClientClassOnce.Do(func() {
		_AVVoiceTriggerClientClass = AVVoiceTriggerClientClass{class: objc.GetClass("AVVoiceTriggerClient")}
	})
	return _AVVoiceTriggerClientClass
}

// GetAVVoiceTriggerClientClass returns the class object for AVVoiceTriggerClient.
func GetAVVoiceTriggerClientClass() AVVoiceTriggerClientClass {
	return getAVVoiceTriggerClientClass()
}

type AVVoiceTriggerClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVoiceTriggerClientClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVoiceTriggerClientClass) Alloc() AVVoiceTriggerClient {
	rv := objc.Send[AVVoiceTriggerClient](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVVoiceTriggerClient.ActivateSecureSession]
//   - [AVVoiceTriggerClient.CallServerCrashedBlock]
//   - [AVVoiceTriggerClient.CallServerResetBlock]
//   - [AVVoiceTriggerClient.ClientType]
//   - [AVVoiceTriggerClient.SetClientType]
//   - [AVVoiceTriggerClient.EnableBargeInModeCompletionBlock]
//   - [AVVoiceTriggerClient.EnableListeningOnPortsCompletionBlock]
//   - [AVVoiceTriggerClient.EnableSpeakerStateListening]
//   - [AVVoiceTriggerClient.EnableSpeakerStateListeningCompletionBlock]
//   - [AVVoiceTriggerClient.EnableVoiceTriggerListening]
//   - [AVVoiceTriggerClient.EnableVoiceTriggerListeningCompletionBlock]
//   - [AVVoiceTriggerClient.GetInputChannelInfoCompletion]
//   - [AVVoiceTriggerClient.IsSiriClient]
//   - [AVVoiceTriggerClient.ListeningEnabledCompletionBlock]
//   - [AVVoiceTriggerClient.PortStateActiveCompletionBlock]
//   - [AVVoiceTriggerClient.PortStateChangedNotification]
//   - [AVVoiceTriggerClient.RecordingAuditTokenList]
//   - [AVVoiceTriggerClient.RecordingPIDList]
//   - [AVVoiceTriggerClient.SecureSessionServerCrash]
//   - [AVVoiceTriggerClient.SecureSessionServerReset]
//   - [AVVoiceTriggerClient.SetAvvcServerCrashedBlock]
//   - [AVVoiceTriggerClient.SetAvvcServerResetBlock]
//   - [AVVoiceTriggerClient.SetAggressiveECModeCompletionBlock]
//   - [AVVoiceTriggerClient.SetListeningPropertyCompletionBlock]
//   - [AVVoiceTriggerClient.SetPortStateChangedBlock]
//   - [AVVoiceTriggerClient.SetServerCrashedBlock]
//   - [AVVoiceTriggerClient.SetServerResetBlock]
//   - [AVVoiceTriggerClient.SetSiriClientRecordStateChangedBlock]
//   - [AVVoiceTriggerClient.SetSpeakerMuteStateChangedBlock]
//   - [AVVoiceTriggerClient.SetSpeakerStateChangedBlock]
//   - [AVVoiceTriggerClient.SetVoiceTriggerBlock]
//   - [AVVoiceTriggerClient.SiriClientRecordStateChangedNotificationRecordingCount]
//   - [AVVoiceTriggerClient.SiriClientsRecordingCompletionBlock]
//   - [AVVoiceTriggerClient.SpeakerMuteStateChangedNotification]
//   - [AVVoiceTriggerClient.SpeakerStateActive]
//   - [AVVoiceTriggerClient.SpeakerStateActiveCompletionBlock]
//   - [AVVoiceTriggerClient.SpeakerStateChangedNotification]
//   - [AVVoiceTriggerClient.SpeakerStateMuted]
//   - [AVVoiceTriggerClient.SpeakerStateMutedCompletionBlock]
//   - [AVVoiceTriggerClient.UpdateVoiceTriggerConfiguration]
//   - [AVVoiceTriggerClient.UpdateVoiceTriggerConfigurationCompletionBlock]
//   - [AVVoiceTriggerClient.VoiceTriggerNotification]
//   - [AVVoiceTriggerClient.VoiceTriggerPastDataFramesAvailable]
//   - [AVVoiceTriggerClient.VoiceTriggerPastDataFramesAvailableCompletion]
//   - [AVVoiceTriggerClient.VoiceTriggerServerConnection]
//   - [AVVoiceTriggerClient.InitWithValue]
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient
type AVVoiceTriggerClient struct {
	objectivec.Object
}

// AVVoiceTriggerClientFromID constructs a [AVVoiceTriggerClient] from an objc.ID.
func AVVoiceTriggerClientFromID(id objc.ID) AVVoiceTriggerClient {
	return AVVoiceTriggerClient{objectivec.Object{ID: id}}
}
// Ensure AVVoiceTriggerClient implements IAVVoiceTriggerClient.
var _ IAVVoiceTriggerClient = AVVoiceTriggerClient{}

// An interface definition for the [AVVoiceTriggerClient] class.
//
// # Methods
//
//   - [IAVVoiceTriggerClient.ActivateSecureSession]
//   - [IAVVoiceTriggerClient.CallServerCrashedBlock]
//   - [IAVVoiceTriggerClient.CallServerResetBlock]
//   - [IAVVoiceTriggerClient.ClientType]
//   - [IAVVoiceTriggerClient.SetClientType]
//   - [IAVVoiceTriggerClient.EnableBargeInModeCompletionBlock]
//   - [IAVVoiceTriggerClient.EnableListeningOnPortsCompletionBlock]
//   - [IAVVoiceTriggerClient.EnableSpeakerStateListening]
//   - [IAVVoiceTriggerClient.EnableSpeakerStateListeningCompletionBlock]
//   - [IAVVoiceTriggerClient.EnableVoiceTriggerListening]
//   - [IAVVoiceTriggerClient.EnableVoiceTriggerListeningCompletionBlock]
//   - [IAVVoiceTriggerClient.GetInputChannelInfoCompletion]
//   - [IAVVoiceTriggerClient.IsSiriClient]
//   - [IAVVoiceTriggerClient.ListeningEnabledCompletionBlock]
//   - [IAVVoiceTriggerClient.PortStateActiveCompletionBlock]
//   - [IAVVoiceTriggerClient.PortStateChangedNotification]
//   - [IAVVoiceTriggerClient.RecordingAuditTokenList]
//   - [IAVVoiceTriggerClient.RecordingPIDList]
//   - [IAVVoiceTriggerClient.SecureSessionServerCrash]
//   - [IAVVoiceTriggerClient.SecureSessionServerReset]
//   - [IAVVoiceTriggerClient.SetAvvcServerCrashedBlock]
//   - [IAVVoiceTriggerClient.SetAvvcServerResetBlock]
//   - [IAVVoiceTriggerClient.SetAggressiveECModeCompletionBlock]
//   - [IAVVoiceTriggerClient.SetListeningPropertyCompletionBlock]
//   - [IAVVoiceTriggerClient.SetPortStateChangedBlock]
//   - [IAVVoiceTriggerClient.SetServerCrashedBlock]
//   - [IAVVoiceTriggerClient.SetServerResetBlock]
//   - [IAVVoiceTriggerClient.SetSiriClientRecordStateChangedBlock]
//   - [IAVVoiceTriggerClient.SetSpeakerMuteStateChangedBlock]
//   - [IAVVoiceTriggerClient.SetSpeakerStateChangedBlock]
//   - [IAVVoiceTriggerClient.SetVoiceTriggerBlock]
//   - [IAVVoiceTriggerClient.SiriClientRecordStateChangedNotificationRecordingCount]
//   - [IAVVoiceTriggerClient.SiriClientsRecordingCompletionBlock]
//   - [IAVVoiceTriggerClient.SpeakerMuteStateChangedNotification]
//   - [IAVVoiceTriggerClient.SpeakerStateActive]
//   - [IAVVoiceTriggerClient.SpeakerStateActiveCompletionBlock]
//   - [IAVVoiceTriggerClient.SpeakerStateChangedNotification]
//   - [IAVVoiceTriggerClient.SpeakerStateMuted]
//   - [IAVVoiceTriggerClient.SpeakerStateMutedCompletionBlock]
//   - [IAVVoiceTriggerClient.UpdateVoiceTriggerConfiguration]
//   - [IAVVoiceTriggerClient.UpdateVoiceTriggerConfigurationCompletionBlock]
//   - [IAVVoiceTriggerClient.VoiceTriggerNotification]
//   - [IAVVoiceTriggerClient.VoiceTriggerPastDataFramesAvailable]
//   - [IAVVoiceTriggerClient.VoiceTriggerPastDataFramesAvailableCompletion]
//   - [IAVVoiceTriggerClient.VoiceTriggerServerConnection]
//   - [IAVVoiceTriggerClient.InitWithValue]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient
type IAVVoiceTriggerClient interface {
	objectivec.IObject

	// Topic: Methods

	ActivateSecureSession(session bool) objectivec.IObject
	CallServerCrashedBlock()
	CallServerResetBlock()
	ClientType() int64
	SetClientType(value int64)
	EnableBargeInModeCompletionBlock(mode bool, block VoidHandler)
	EnableListeningOnPortsCompletionBlock(ports objectivec.IObject, block VoidHandler)
	EnableSpeakerStateListening(listening bool)
	EnableSpeakerStateListeningCompletionBlock(listening bool, block VoidHandler)
	EnableVoiceTriggerListening(listening bool)
	EnableVoiceTriggerListeningCompletionBlock(listening bool, block VoidHandler)
	GetInputChannelInfoCompletion(completion VoidHandler)
	IsSiriClient() bool
	ListeningEnabledCompletionBlock(block VoidHandler)
	PortStateActiveCompletionBlock(block VoidHandler)
	PortStateChangedNotification(notification objectivec.IObject)
	RecordingAuditTokenList() objectivec.IObject
	RecordingPIDList() objectivec.IObject
	SecureSessionServerCrash()
	SecureSessionServerReset()
	SetAvvcServerCrashedBlock(block VoidHandler)
	SetAvvcServerResetBlock(block VoidHandler)
	SetAggressiveECModeCompletionBlock(eCMode bool, block VoidHandler)
	SetListeningPropertyCompletionBlock(property bool, block VoidHandler)
	SetPortStateChangedBlock(block VoidHandler)
	SetServerCrashedBlock(block VoidHandler)
	SetServerResetBlock(block VoidHandler)
	SetSiriClientRecordStateChangedBlock(block VoidHandler)
	SetSpeakerMuteStateChangedBlock(block VoidHandler)
	SetSpeakerStateChangedBlock(block VoidHandler)
	SetVoiceTriggerBlock(block VoidHandler)
	SiriClientRecordStateChangedNotificationRecordingCount(notification bool, count uint64)
	SiriClientsRecordingCompletionBlock(block VoidHandler)
	SpeakerMuteStateChangedNotification(notification bool)
	SpeakerStateActive() bool
	SpeakerStateActiveCompletionBlock(block VoidHandler)
	SpeakerStateChangedNotification(notification objectivec.IObject)
	SpeakerStateMuted() bool
	SpeakerStateMutedCompletionBlock(block VoidHandler)
	UpdateVoiceTriggerConfiguration(configuration objectivec.IObject)
	UpdateVoiceTriggerConfigurationCompletionBlock(configuration objectivec.IObject, block VoidHandler)
	VoiceTriggerNotification(notification objectivec.IObject)
	VoiceTriggerPastDataFramesAvailable() uint64
	VoiceTriggerPastDataFramesAvailableCompletion(completion VoidHandler)
	VoiceTriggerServerConnection() objectivec.IObject
	InitWithValue(init_ int64) AVVoiceTriggerClient
}

// Init initializes the instance.
func (v AVVoiceTriggerClient) Init() AVVoiceTriggerClient {
	rv := objc.Send[AVVoiceTriggerClient](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVoiceTriggerClient) Autorelease() AVVoiceTriggerClient {
	rv := objc.Send[AVVoiceTriggerClient](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVoiceTriggerClient creates a new AVVoiceTriggerClient instance.
func NewAVVoiceTriggerClient() AVVoiceTriggerClient {
	class := getAVVoiceTriggerClientClass()
	rv := objc.Send[AVVoiceTriggerClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/init:
func NewVoiceTriggerClientWithValue(init_ int64) AVVoiceTriggerClient {
	instance := getAVVoiceTriggerClientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("init:"), init_)
	return AVVoiceTriggerClientFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/activateSecureSession:
func (v AVVoiceTriggerClient) ActivateSecureSession(session bool) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("activateSecureSession:"), session)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/callServerCrashedBlock
func (v AVVoiceTriggerClient) CallServerCrashedBlock() {
	objc.Send[objc.ID](v.ID, objc.Sel("callServerCrashedBlock"))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/callServerResetBlock
func (v AVVoiceTriggerClient) CallServerResetBlock() {
	objc.Send[objc.ID](v.ID, objc.Sel("callServerResetBlock"))
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/enableBargeInMode:completionBlock:
func (v AVVoiceTriggerClient) EnableBargeInModeCompletionBlock(mode bool, block VoidHandler) {
_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("enableBargeInMode:completionBlock:"), mode, _block1)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/enableListeningOnPorts:completionBlock:
func (v AVVoiceTriggerClient) EnableListeningOnPortsCompletionBlock(ports objectivec.IObject, block VoidHandler) {
_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("enableListeningOnPorts:completionBlock:"), ports, _block1)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/enableSpeakerStateListening:
func (v AVVoiceTriggerClient) EnableSpeakerStateListening(listening bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("enableSpeakerStateListening:"), listening)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/enableSpeakerStateListening:completionBlock:
func (v AVVoiceTriggerClient) EnableSpeakerStateListeningCompletionBlock(listening bool, block VoidHandler) {
_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("enableSpeakerStateListening:completionBlock:"), listening, _block1)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/enableVoiceTriggerListening:
func (v AVVoiceTriggerClient) EnableVoiceTriggerListening(listening bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("enableVoiceTriggerListening:"), listening)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/enableVoiceTriggerListening:completionBlock:
func (v AVVoiceTriggerClient) EnableVoiceTriggerListeningCompletionBlock(listening bool, block VoidHandler) {
_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("enableVoiceTriggerListening:completionBlock:"), listening, _block1)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/getInputChannelInfoCompletion:
func (v AVVoiceTriggerClient) GetInputChannelInfoCompletion(completion VoidHandler) {
_block0, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("getInputChannelInfoCompletion:"), _block0)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/isSiriClient
func (v AVVoiceTriggerClient) IsSiriClient() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isSiriClient"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/listeningEnabledCompletionBlock:
func (v AVVoiceTriggerClient) ListeningEnabledCompletionBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("listeningEnabledCompletionBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/portStateActiveCompletionBlock:
func (v AVVoiceTriggerClient) PortStateActiveCompletionBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("portStateActiveCompletionBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/portStateChangedNotification:
func (v AVVoiceTriggerClient) PortStateChangedNotification(notification objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("portStateChangedNotification:"), notification)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/recordingAuditTokenList
func (v AVVoiceTriggerClient) RecordingAuditTokenList() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("recordingAuditTokenList"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/recordingPIDList
func (v AVVoiceTriggerClient) RecordingPIDList() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("recordingPIDList"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/secureSessionServerCrash
func (v AVVoiceTriggerClient) SecureSessionServerCrash() {
	objc.Send[objc.ID](v.ID, objc.Sel("secureSessionServerCrash"))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/secureSessionServerReset
func (v AVVoiceTriggerClient) SecureSessionServerReset() {
	objc.Send[objc.ID](v.ID, objc.Sel("secureSessionServerReset"))
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setAggressiveECMode:completionBlock:
func (v AVVoiceTriggerClient) SetAggressiveECModeCompletionBlock(eCMode bool, block VoidHandler) {
_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setAggressiveECMode:completionBlock:"), eCMode, _block1)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setAvvcServerCrashedBlock:
func (v AVVoiceTriggerClient) SetAvvcServerCrashedBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setAvvcServerCrashedBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setAvvcServerResetBlock:
func (v AVVoiceTriggerClient) SetAvvcServerResetBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setAvvcServerResetBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setListeningProperty:completionBlock:
func (v AVVoiceTriggerClient) SetListeningPropertyCompletionBlock(property bool, block VoidHandler) {
_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setListeningProperty:completionBlock:"), property, _block1)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setPortStateChangedBlock:
func (v AVVoiceTriggerClient) SetPortStateChangedBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setPortStateChangedBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setServerCrashedBlock:
func (v AVVoiceTriggerClient) SetServerCrashedBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setServerCrashedBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setServerResetBlock:
func (v AVVoiceTriggerClient) SetServerResetBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setServerResetBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setSiriClientRecordStateChangedBlock:
func (v AVVoiceTriggerClient) SetSiriClientRecordStateChangedBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setSiriClientRecordStateChangedBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setSpeakerMuteStateChangedBlock:
func (v AVVoiceTriggerClient) SetSpeakerMuteStateChangedBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setSpeakerMuteStateChangedBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setSpeakerStateChangedBlock:
func (v AVVoiceTriggerClient) SetSpeakerStateChangedBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setSpeakerStateChangedBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/setVoiceTriggerBlock:
func (v AVVoiceTriggerClient) SetVoiceTriggerBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setVoiceTriggerBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/siriClientRecordStateChangedNotification:recordingCount:
func (v AVVoiceTriggerClient) SiriClientRecordStateChangedNotificationRecordingCount(notification bool, count uint64) {
	objc.Send[objc.ID](v.ID, objc.Sel("siriClientRecordStateChangedNotification:recordingCount:"), notification, count)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/siriClientsRecordingCompletionBlock:
func (v AVVoiceTriggerClient) SiriClientsRecordingCompletionBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("siriClientsRecordingCompletionBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/speakerMuteStateChangedNotification:
func (v AVVoiceTriggerClient) SpeakerMuteStateChangedNotification(notification bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("speakerMuteStateChangedNotification:"), notification)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/speakerStateActive
func (v AVVoiceTriggerClient) SpeakerStateActive() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("speakerStateActive"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/speakerStateActiveCompletionBlock:
func (v AVVoiceTriggerClient) SpeakerStateActiveCompletionBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("speakerStateActiveCompletionBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/speakerStateChangedNotification:
func (v AVVoiceTriggerClient) SpeakerStateChangedNotification(notification objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("speakerStateChangedNotification:"), notification)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/speakerStateMuted
func (v AVVoiceTriggerClient) SpeakerStateMuted() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("speakerStateMuted"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/speakerStateMutedCompletionBlock:
func (v AVVoiceTriggerClient) SpeakerStateMutedCompletionBlock(block VoidHandler) {
_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("speakerStateMutedCompletionBlock:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/updateVoiceTriggerConfiguration:
func (v AVVoiceTriggerClient) UpdateVoiceTriggerConfiguration(configuration objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("updateVoiceTriggerConfiguration:"), configuration)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/updateVoiceTriggerConfiguration:completionBlock:
func (v AVVoiceTriggerClient) UpdateVoiceTriggerConfigurationCompletionBlock(configuration objectivec.IObject, block VoidHandler) {
_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("updateVoiceTriggerConfiguration:completionBlock:"), configuration, _block1)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/voiceTriggerNotification:
func (v AVVoiceTriggerClient) VoiceTriggerNotification(notification objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("voiceTriggerNotification:"), notification)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/voiceTriggerPastDataFramesAvailableCompletion:
func (v AVVoiceTriggerClient) VoiceTriggerPastDataFramesAvailableCompletion(completion VoidHandler) {
_block0, _ := NewVoidBlock(completion)
	objc.Send[objc.ID](v.ID, objc.Sel("voiceTriggerPastDataFramesAvailableCompletion:"), _block0)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/voiceTriggerServerConnection
func (v AVVoiceTriggerClient) VoiceTriggerServerConnection() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("voiceTriggerServerConnection"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/init:
func (v AVVoiceTriggerClient) InitWithValue(init_ int64) AVVoiceTriggerClient {
	rv := objc.Send[AVVoiceTriggerClient](v.ID, objc.Sel("init:"), init_)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/isAPIAvailable
func (_AVVoiceTriggerClientClass AVVoiceTriggerClientClass) IsAPIAvailable() bool {
	rv := objc.Send[bool](objc.ID(_AVVoiceTriggerClientClass.class), objc.Sel("isAPIAvailable"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/sharedInstance
func (_AVVoiceTriggerClientClass AVVoiceTriggerClientClass) SharedInstance() AVVoiceTriggerClient {
	rv := objc.Send[objc.ID](objc.ID(_AVVoiceTriggerClientClass.class), objc.Sel("sharedInstance"))
	return AVVoiceTriggerClientFromID(rv)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/sharedInstance:
func (_AVVoiceTriggerClientClass AVVoiceTriggerClientClass) SharedInstanceWithInstance(instance int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_AVVoiceTriggerClientClass.class), objc.Sel("sharedInstance:"), instance)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/supportsDuckingOnSpeakerOutput
func (_AVVoiceTriggerClientClass AVVoiceTriggerClientClass) SupportsDuckingOnSpeakerOutput() bool {
	rv := objc.Send[bool](objc.ID(_AVVoiceTriggerClientClass.class), objc.Sel("supportsDuckingOnSpeakerOutput"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/clientType
func (v AVVoiceTriggerClient) ClientType() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("clientType"))
	return rv
}
func (v AVVoiceTriggerClient) SetClientType(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setClientType:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClient/voiceTriggerPastDataFramesAvailable
func (v AVVoiceTriggerClient) VoiceTriggerPastDataFramesAvailable() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("voiceTriggerPastDataFramesAvailable"))
	return rv
}

// EnableBargeInModeCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.EnableBargeInModeCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) EnableBargeInModeCompletionBlockSync(ctx context.Context, mode bool) error {
	done := make(chan struct{}, 1)
	v.EnableBargeInModeCompletionBlock(mode, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnableListeningOnPortsCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.EnableListeningOnPortsCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) EnableListeningOnPortsCompletionBlockSync(ctx context.Context, ports objectivec.IObject) error {
	done := make(chan struct{}, 1)
	v.EnableListeningOnPortsCompletionBlock(ports, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnableSpeakerStateListeningCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.EnableSpeakerStateListeningCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) EnableSpeakerStateListeningCompletionBlockSync(ctx context.Context, listening bool) error {
	done := make(chan struct{}, 1)
	v.EnableSpeakerStateListeningCompletionBlock(listening, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EnableVoiceTriggerListeningCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.EnableVoiceTriggerListeningCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) EnableVoiceTriggerListeningCompletionBlockSync(ctx context.Context, listening bool) error {
	done := make(chan struct{}, 1)
	v.EnableVoiceTriggerListeningCompletionBlock(listening, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetInputChannelInfoCompletionSync is a synchronous wrapper around [AVVoiceTriggerClient.GetInputChannelInfoCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) GetInputChannelInfoCompletionSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.GetInputChannelInfoCompletion(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ListeningEnabledCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.ListeningEnabledCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) ListeningEnabledCompletionBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.ListeningEnabledCompletionBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PortStateActiveCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.PortStateActiveCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) PortStateActiveCompletionBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.PortStateActiveCompletionBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetAggressiveECModeCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetAggressiveECModeCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetAggressiveECModeCompletionBlockSync(ctx context.Context, eCMode bool) error {
	done := make(chan struct{}, 1)
	v.SetAggressiveECModeCompletionBlock(eCMode, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetAvvcServerCrashedBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetAvvcServerCrashedBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetAvvcServerCrashedBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetAvvcServerCrashedBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetAvvcServerResetBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetAvvcServerResetBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetAvvcServerResetBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetAvvcServerResetBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetListeningPropertyCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetListeningPropertyCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetListeningPropertyCompletionBlockSync(ctx context.Context, property bool) error {
	done := make(chan struct{}, 1)
	v.SetListeningPropertyCompletionBlock(property, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetPortStateChangedBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetPortStateChangedBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetPortStateChangedBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetPortStateChangedBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetServerCrashedBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetServerCrashedBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetServerCrashedBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetServerCrashedBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetServerResetBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetServerResetBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetServerResetBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetServerResetBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetSiriClientRecordStateChangedBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetSiriClientRecordStateChangedBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetSiriClientRecordStateChangedBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetSiriClientRecordStateChangedBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetSpeakerMuteStateChangedBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetSpeakerMuteStateChangedBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetSpeakerMuteStateChangedBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetSpeakerMuteStateChangedBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetSpeakerStateChangedBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetSpeakerStateChangedBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetSpeakerStateChangedBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetSpeakerStateChangedBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetVoiceTriggerBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SetVoiceTriggerBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SetVoiceTriggerBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetVoiceTriggerBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SiriClientsRecordingCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SiriClientsRecordingCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SiriClientsRecordingCompletionBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SiriClientsRecordingCompletionBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SpeakerStateActiveCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SpeakerStateActiveCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SpeakerStateActiveCompletionBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SpeakerStateActiveCompletionBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SpeakerStateMutedCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.SpeakerStateMutedCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) SpeakerStateMutedCompletionBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SpeakerStateMutedCompletionBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// UpdateVoiceTriggerConfigurationCompletionBlockSync is a synchronous wrapper around [AVVoiceTriggerClient.UpdateVoiceTriggerConfigurationCompletionBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) UpdateVoiceTriggerConfigurationCompletionBlockSync(ctx context.Context, configuration objectivec.IObject) error {
	done := make(chan struct{}, 1)
	v.UpdateVoiceTriggerConfigurationCompletionBlock(configuration, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// VoiceTriggerPastDataFramesAvailableCompletionSync is a synchronous wrapper around [AVVoiceTriggerClient.VoiceTriggerPastDataFramesAvailableCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClient) VoiceTriggerPastDataFramesAvailableCompletionSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.VoiceTriggerPastDataFramesAvailableCompletion(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

