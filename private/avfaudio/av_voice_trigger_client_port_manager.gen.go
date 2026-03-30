// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVoiceTriggerClientPortManager] class.
var (
	_AVVoiceTriggerClientPortManagerClass     AVVoiceTriggerClientPortManagerClass
	_AVVoiceTriggerClientPortManagerClassOnce sync.Once
)

func getAVVoiceTriggerClientPortManagerClass() AVVoiceTriggerClientPortManagerClass {
	_AVVoiceTriggerClientPortManagerClassOnce.Do(func() {
		_AVVoiceTriggerClientPortManagerClass = AVVoiceTriggerClientPortManagerClass{class: objc.GetClass("AVVoiceTriggerClientPortManager")}
	})
	return _AVVoiceTriggerClientPortManagerClass
}

// GetAVVoiceTriggerClientPortManagerClass returns the class object for AVVoiceTriggerClientPortManager.
func GetAVVoiceTriggerClientPortManagerClass() AVVoiceTriggerClientPortManagerClass {
	return getAVVoiceTriggerClientPortManagerClass()
}

type AVVoiceTriggerClientPortManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVoiceTriggerClientPortManagerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVoiceTriggerClientPortManagerClass) Alloc() AVVoiceTriggerClientPortManager {
	rv := objc.Send[AVVoiceTriggerClientPortManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVoiceTriggerClientPortManager.CallMuteStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.CallRunningStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.DeviceID]
//   - [AVVoiceTriggerClientPortManager.SetDeviceID]
//   - [AVVoiceTriggerClientPortManager.Generation]
//   - [AVVoiceTriggerClientPortManager.SetGeneration]
//   - [AVVoiceTriggerClientPortManager.HysteresisDurationSeconds]
//   - [AVVoiceTriggerClientPortManager.SetHysteresisDurationSeconds]
//   - [AVVoiceTriggerClientPortManager.LastRunningStateSent]
//   - [AVVoiceTriggerClientPortManager.SetLastRunningStateSent]
//   - [AVVoiceTriggerClientPortManager.ListeningEnabled]
//   - [AVVoiceTriggerClientPortManager.SetListeningEnabled]
//   - [AVVoiceTriggerClientPortManager.MuteStateChangeNotificationRegistered]
//   - [AVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationRegistered]
//   - [AVVoiceTriggerClientPortManager.NotifyMuteStateChanged]
//   - [AVVoiceTriggerClientPortManager.NotifyRunningStateChangedWithHysteresis]
//   - [AVVoiceTriggerClientPortManager.PortType]
//   - [AVVoiceTriggerClientPortManager.SetPortType]
//   - [AVVoiceTriggerClientPortManager.Queue]
//   - [AVVoiceTriggerClientPortManager.SetQueue]
//   - [AVVoiceTriggerClientPortManager.RunningStateChangeNotificationRegistered]
//   - [AVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationRegistered]
//   - [AVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationBlock]
//   - [AVVoiceTriggerClientPortManager.InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager
type AVVoiceTriggerClientPortManager struct {
	objectivec.Object
}

// AVVoiceTriggerClientPortManagerFromID constructs a [AVVoiceTriggerClientPortManager] from an objc.ID.
func AVVoiceTriggerClientPortManagerFromID(id objc.ID) AVVoiceTriggerClientPortManager {
	return AVVoiceTriggerClientPortManager{objectivec.Object{ID: id}}
}

// Ensure AVVoiceTriggerClientPortManager implements IAVVoiceTriggerClientPortManager.
var _ IAVVoiceTriggerClientPortManager = AVVoiceTriggerClientPortManager{}

// An interface definition for the [AVVoiceTriggerClientPortManager] class.
//
// # Methods
//
//   - [IAVVoiceTriggerClientPortManager.CallMuteStateChangeNotificationBlock]
//   - [IAVVoiceTriggerClientPortManager.CallRunningStateChangeNotificationBlock]
//   - [IAVVoiceTriggerClientPortManager.DeviceID]
//   - [IAVVoiceTriggerClientPortManager.SetDeviceID]
//   - [IAVVoiceTriggerClientPortManager.Generation]
//   - [IAVVoiceTriggerClientPortManager.SetGeneration]
//   - [IAVVoiceTriggerClientPortManager.HysteresisDurationSeconds]
//   - [IAVVoiceTriggerClientPortManager.SetHysteresisDurationSeconds]
//   - [IAVVoiceTriggerClientPortManager.LastRunningStateSent]
//   - [IAVVoiceTriggerClientPortManager.SetLastRunningStateSent]
//   - [IAVVoiceTriggerClientPortManager.ListeningEnabled]
//   - [IAVVoiceTriggerClientPortManager.SetListeningEnabled]
//   - [IAVVoiceTriggerClientPortManager.MuteStateChangeNotificationRegistered]
//   - [IAVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationRegistered]
//   - [IAVVoiceTriggerClientPortManager.NotifyMuteStateChanged]
//   - [IAVVoiceTriggerClientPortManager.NotifyRunningStateChangedWithHysteresis]
//   - [IAVVoiceTriggerClientPortManager.PortType]
//   - [IAVVoiceTriggerClientPortManager.SetPortType]
//   - [IAVVoiceTriggerClientPortManager.Queue]
//   - [IAVVoiceTriggerClientPortManager.SetQueue]
//   - [IAVVoiceTriggerClientPortManager.RunningStateChangeNotificationRegistered]
//   - [IAVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationRegistered]
//   - [IAVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationBlock]
//   - [IAVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationBlock]
//   - [IAVVoiceTriggerClientPortManager.InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager
type IAVVoiceTriggerClientPortManager interface {
	objectivec.IObject

	// Topic: Methods

	CallMuteStateChangeNotificationBlock(block bool)
	CallRunningStateChangeNotificationBlock(block bool)
	DeviceID() uint32
	SetDeviceID(value uint32)
	Generation() int64
	SetGeneration(value int64)
	HysteresisDurationSeconds() float32
	SetHysteresisDurationSeconds(value float32)
	LastRunningStateSent() bool
	SetLastRunningStateSent(value bool)
	ListeningEnabled() bool
	SetListeningEnabled(value bool)
	MuteStateChangeNotificationRegistered() bool
	SetMuteStateChangeNotificationRegistered(value bool)
	NotifyMuteStateChanged()
	NotifyRunningStateChangedWithHysteresis()
	PortType() uint64
	SetPortType(value uint64)
	Queue() objectivec.Object
	SetQueue(value objectivec.Object)
	RunningStateChangeNotificationRegistered() bool
	SetRunningStateChangeNotificationRegistered(value bool)
	SetMuteStateChangeNotificationBlock(block VoidHandler)
	SetRunningStateChangeNotificationBlock(block VoidHandler)
	InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock(queue objectivec.IObject, type_ uint64, seconds float32, block VoidHandler, block2 VoidHandler) AVVoiceTriggerClientPortManager
}

// Init initializes the instance.
func (v AVVoiceTriggerClientPortManager) Init() AVVoiceTriggerClientPortManager {
	rv := objc.Send[AVVoiceTriggerClientPortManager](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVoiceTriggerClientPortManager) Autorelease() AVVoiceTriggerClientPortManager {
	rv := objc.Send[AVVoiceTriggerClientPortManager](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVoiceTriggerClientPortManager creates a new AVVoiceTriggerClientPortManager instance.
func NewAVVoiceTriggerClientPortManager() AVVoiceTriggerClientPortManager {
	class := getAVVoiceTriggerClientPortManagerClass()
	rv := objc.Send[AVVoiceTriggerClientPortManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/callMuteStateChangeNotificationBlock:
func (v AVVoiceTriggerClientPortManager) CallMuteStateChangeNotificationBlock(block bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("callMuteStateChangeNotificationBlock:"), block)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/callRunningStateChangeNotificationBlock:
func (v AVVoiceTriggerClientPortManager) CallRunningStateChangeNotificationBlock(block bool) {
	objc.Send[objc.ID](v.ID, objc.Sel("callRunningStateChangeNotificationBlock:"), block)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/notifyMuteStateChanged
func (v AVVoiceTriggerClientPortManager) NotifyMuteStateChanged() {
	objc.Send[objc.ID](v.ID, objc.Sel("notifyMuteStateChanged"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/notifyRunningStateChangedWithHysteresis
func (v AVVoiceTriggerClientPortManager) NotifyRunningStateChangedWithHysteresis() {
	objc.Send[objc.ID](v.ID, objc.Sel("notifyRunningStateChangedWithHysteresis"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/setMuteStateChangeNotificationBlock:
func (v AVVoiceTriggerClientPortManager) SetMuteStateChangeNotificationBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setMuteStateChangeNotificationBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/setRunningStateChangeNotificationBlock:
func (v AVVoiceTriggerClientPortManager) SetRunningStateChangeNotificationBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.Send[objc.ID](v.ID, objc.Sel("setRunningStateChangeNotificationBlock:"), _block0)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/initWithSerialQueue:portType:hysteresisDurationSeconds:runningStateChangeNotificationBlock:muteStateChangeNotificationBlock:
func (v AVVoiceTriggerClientPortManager) InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock(queue objectivec.IObject, type_ uint64, seconds float32, block VoidHandler, block2 VoidHandler) AVVoiceTriggerClientPortManager {
	_block3, _ := NewVoidBlock(block)
	_block4, _ := NewVoidBlock(block2)
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithSerialQueue:portType:hysteresisDurationSeconds:runningStateChangeNotificationBlock:muteStateChangeNotificationBlock:"), queue, type_, seconds, _block3, _block4)
	return AVVoiceTriggerClientPortManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/deviceID
func (v AVVoiceTriggerClientPortManager) DeviceID() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("deviceID"))
	return rv
}
func (v AVVoiceTriggerClientPortManager) SetDeviceID(value uint32) {
	objc.Send[struct{}](v.ID, objc.Sel("setDeviceID:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/generation
func (v AVVoiceTriggerClientPortManager) Generation() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("generation"))
	return rv
}
func (v AVVoiceTriggerClientPortManager) SetGeneration(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setGeneration:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/hysteresisDurationSeconds
func (v AVVoiceTriggerClientPortManager) HysteresisDurationSeconds() float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("hysteresisDurationSeconds"))
	return rv
}
func (v AVVoiceTriggerClientPortManager) SetHysteresisDurationSeconds(value float32) {
	objc.Send[struct{}](v.ID, objc.Sel("setHysteresisDurationSeconds:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/lastRunningStateSent
func (v AVVoiceTriggerClientPortManager) LastRunningStateSent() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("lastRunningStateSent"))
	return rv
}
func (v AVVoiceTriggerClientPortManager) SetLastRunningStateSent(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setLastRunningStateSent:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/listeningEnabled
func (v AVVoiceTriggerClientPortManager) ListeningEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("listeningEnabled"))
	return rv
}
func (v AVVoiceTriggerClientPortManager) SetListeningEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setListeningEnabled:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/muteStateChangeNotificationRegistered
func (v AVVoiceTriggerClientPortManager) MuteStateChangeNotificationRegistered() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("muteStateChangeNotificationRegistered"))
	return rv
}
func (v AVVoiceTriggerClientPortManager) SetMuteStateChangeNotificationRegistered(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setMuteStateChangeNotificationRegistered:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/portType
func (v AVVoiceTriggerClientPortManager) PortType() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("portType"))
	return rv
}
func (v AVVoiceTriggerClientPortManager) SetPortType(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setPortType:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/queue
func (v AVVoiceTriggerClientPortManager) Queue() objectivec.Object {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (v AVVoiceTriggerClientPortManager) SetQueue(value objectivec.Object) {
	objc.Send[struct{}](v.ID, objc.Sel("setQueue:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVoiceTriggerClientPortManager/runningStateChangeNotificationRegistered
func (v AVVoiceTriggerClientPortManager) RunningStateChangeNotificationRegistered() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("runningStateChangeNotificationRegistered"))
	return rv
}
func (v AVVoiceTriggerClientPortManager) SetRunningStateChangeNotificationRegistered(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setRunningStateChangeNotificationRegistered:"), value)
}

// SetMuteStateChangeNotificationBlockSync is a synchronous wrapper around [AVVoiceTriggerClientPortManager.SetMuteStateChangeNotificationBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClientPortManager) SetMuteStateChangeNotificationBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetMuteStateChangeNotificationBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetRunningStateChangeNotificationBlockSync is a synchronous wrapper around [AVVoiceTriggerClientPortManager.SetRunningStateChangeNotificationBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClientPortManager) SetRunningStateChangeNotificationBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	v.SetRunningStateChangeNotificationBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlockSync is a synchronous wrapper around [AVVoiceTriggerClientPortManager.InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (v AVVoiceTriggerClientPortManager) InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlockSync(ctx context.Context, queue objectivec.IObject, type_ uint64, seconds float32, block VoidHandler) error {
	done := make(chan struct{}, 1)
	v.InitWithSerialQueuePortTypeHysteresisDurationSecondsRunningStateChangeNotificationBlockMuteStateChangeNotificationBlock(queue, type_, seconds, block, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
